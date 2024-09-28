package main

import (
	"fmt"
	"github.com/Jumpaku/eyamluate/golang/eval"
	yaml2 "github.com/Jumpaku/eyamluate/golang/yaml"
	"github.com/goccy/go-yaml"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"log"
	"os"
)

func panicOnInputError(inErr error, subcmd []string) {
	if inErr != nil {
		fmt.Fprintln(os.Stderr, GetDoc(subcmd))
		log.Panicln(inErr)
	}
}
func panicOnError(err error, format string, args ...any) {
	if err != nil {
		log.Panicln(fmt.Errorf(format+": %w", append(append([]any{}, args...), err)))
	}
}
func panicIf(panicCond bool, format string, args ...any) {
	if panicCond {
		log.Panicln(fmt.Errorf(format, args...))
	}
}
func readerOrStdin(inFile string) io.ReadCloser {
	if inFile == "" {
		return io.NopCloser(os.Stdin)
	}
	f, err := os.Open(inFile)
	panicOnError(err, "fail to open file", inFile)
	return f
}
func writerOrStdout(outFile string) io.WriteCloser {
	if outFile == "" {
		return os.Stdout
	}
	f, err := os.Create(outFile)
	panicOnError(err, "fail to create file", outFile)
	return f

}
func exitAfterHelp(flag bool, subcmd []string) {
	if flag {
		fmt.Println(GetDoc(subcmd))
		os.Exit(0)
	}
}

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcmd []string, in CLI_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)

		fmt.Println(GetDoc(subcmd))
		return nil
	}
	cli.Version.FUNC = func(subcmd []string, in CLI_Version_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		fmt.Println("v0.0.0")
		return nil
	}
	cli.Validate.FUNC = func(subcmd []string, in CLI_Validate_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)
		inFile := readerOrStdin(in.Opt_InputPath)
		defer inFile.Close()
		outFile := writerOrStdout(in.Opt_OutputPath)
		defer outFile.Close()

		v := eval.NewValidator()
		s, err := io.ReadAll(inFile)
		panicOnError(err, "fail to read file")
		result := v.Validate(&eval.ValidateInput{Source: string(s)})

		_, err = io.WriteString(outFile, result.Status.String()+"\n")
		panicOnError(err, "fail to write file")
		_, err = io.WriteString(outFile, result.ErrorMessage+"\n")
		panicOnError(err, "fail to write file")

		return nil
	}
	cli.Eval.FUNC = func(subcmd []string, in CLI_Eval_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)

		var format yaml2.EncodeFormat
		switch in.Opt_Format {
		case "json":
			format = yaml2.EncodeFormat_ENCODE_FORMAT_JSON
		case "yaml":
			format = yaml2.EncodeFormat_ENCODE_FORMAT_YAML
		default:
			log.Panicf("format must be 'json' or 'yaml': %q\n", in.Opt_Format)
		}

		inFile := readerOrStdin(in.Opt_InputPath)
		defer inFile.Close()
		outFile := writerOrStdout(in.Opt_OutputPath)
		defer outFile.Close()

		e := eval.NewEvaluator()
		s, err := io.ReadAll(inFile)
		panicOnError(err, "fail to read file")
		result := e.Evaluate(&eval.EvaluateInput{Source: string(s)})
		if result.Status != eval.EvaluateOutput_OK {
			m := protojson.MarshalOptions{}
			if in.Opt_Pretty {
				m.Indent = "  "
			}
			b, err := m.Marshal(result)
			panicOnError(err, "fail to marshal json")

			switch format {
			case yaml2.EncodeFormat_ENCODE_FORMAT_JSON:
				_, err = io.WriteString(outFile, string(b))
				panicOnError(err, "fail to write file")
			case yaml2.EncodeFormat_ENCODE_FORMAT_YAML:
				var v any
				err := yaml.Unmarshal(b, &v)
				panicOnError(err, "fail to unmarshal json")
				e := yaml.NewEncoder(outFile, yaml.UseLiteralStyleIfMultiline(true))
				if !in.Opt_Pretty {
					e = yaml.NewEncoder(outFile, yaml.Flow(true))
				}
				err = e.Encode(v)
				panicOnError(err, "fail to marshal yaml")
			}
			return nil
		}
		v := yaml2.NewEncoder().Encode(&yaml2.EncodeInput{
			Format: format,
			Pretty: in.Opt_Pretty,
			Value:  result.Value,
		})
		panicIf(v.IsError, "fail to encode value: %s", v.ErrorMessage)
		_, err = io.WriteString(outFile, v.Result)
		panicOnError(err, "fail to write file")
		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}
