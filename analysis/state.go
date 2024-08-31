package analysis

import (
	"educationalsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {

		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 1,
				Source:   "Common Sense",
				Message:  "Great choice :)",
			})

		}
		if strings.Contains(line, "Rcb") {
			idx := strings.Index(line, "Rcb")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Rcb")),
				Severity: 4,
				Source:   "Common Sense",
				Message:  "Great choice E Sala Cup Namdhe :)",
			})
		}

		if strings.Contains(line, "puts") {
			idx := strings.Index(line, "puts")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("puts")),
				Severity: 4,
				Source:   "tcl doc",
				Message:  "it displays message ",
			})

		}
		if strings.Contains(line, "prathvi") {
			idx := strings.Index(line, "prathvi")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("prathvi")),
				Severity: 4,
				Source:   "tcl doc",
				Message:  "ise , b sec ",
			})

		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// In real life, this would look up the type in our type analysis code...

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// In real life, this would look up the definition

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "Neovim")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("Neovim")),
					NewText: "VS Code",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace Neovim with VS Code",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("Neovim")),
					NewText: "Neo*im",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to Neo*im",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {

	// Ask your static analysis tools to figure out good completions
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "Fun to watch in videos. Don't forget to like & subscribe to streamers using it :)",
		},
		{
			Label:         "Preethamgowda",
			Detail:        "ise b sec",
			Documentation: "acharya institute of technology",
		},
		{
			Label:         "puts ",
			Detail:        "puts ?-nonewline? ?channelId? string",
			Documentation: "used to display message",
		},
		{
			Label:         "after ",
			Detail:        " Execute a command after a time delay",
			Documentation: "Write to a channel",
		},
		{
			Label:         "append ",
			Detail:        " append varName ?value value value ...?",
			Documentation: "Append to variable",
		},
		{
			Label:         "array ",
			Detail:        " Manipulate array variables",
			Documentation: "operation on array can be performed",
		},
		{
			Label:         "binary ",
			Detail:        "  Insert and extract fields from binary strings",
			Documentation: "operation can be performed",
		},
		{
			Label:         "break ",
			Detail:        "Abort looping command",
			Documentation: "",
		},
		{
			Label:         "cd ",
			Detail:        "cd ?dirName?",
			Documentation: "Change working directory",
		},
		{
			Label:         "close ",
			Detail:        "close channelId ?r(ead)|w(rite)?",
			Documentation: "Close an open channel",
		},
		{
			Label:         "concat",
			Detail:        "concat ?arg arg ...?",
			Documentation: "Join lists together",
		},
		{
			Label:         "const",
			Detail:        "const varName value",
			Documentation: " create and initialize a constant",
		},
		{
			Label:         "continue",
			Detail:        "continue",
			Documentation: "Skip to the next iteration of a loop",
		},
		{
			Label:         "eof",
			Detail:        "eof channelId",
			Documentation: "Check for end of file condition on channel",
		},
		{
			Label:         "error",
			Detail:        "error message ?info? ?code?",
			Documentation: "Generate an error",
		},
		{
			Label:         "eval",
			Detail:        "eval arg ?arg ...?",
			Documentation: " Evaluate a Tcl script",
		},
		{
			Label:         "exit ",
			Detail:        "exit ?returnCode?",
			Documentation: "End the application",
		},
		{
			Label:         "chan ",
			Detail:        "chan operation ?arg arg ...?",
			Documentation: "Reads, writes and manipulates channels.",
		},
		{
			Label:         "fcopy",
			Detail:        "fcopy inputChan outputChan ?-size size? ?-command callback?",
			Documentation: " Copy data from one channel to another",
		},
		{
			Label:         "file ",
			Detail:        "file option name ?arg arg ...?",
			Documentation: "Manipulate file names and attributes",
		},
		{
			Label:         "for",
			Detail:        "for start test next body",
			Documentation: " 'For' loop",
		},
		{
			Label:         "gets",
			Detail:        "gets channelId ?varName?",
			Documentation: " Read a line from a channel",
		},
		{
			Label:         "global ",
			Detail:        "global ?varname ...?",
			Documentation: " Access global variables",
		},
		{
			Label:         "http ",
			Detail:        "",
			Documentation: " Access global variablesClient-side implementation of the HTTP/1.1 protocol",
		},
		{
			Label:         "if",
			Detail:        "if expr1 ?then? body1 elseif expr2 ?then? body2 elseif ... ?else? ?bodyN?",
			Documentation: "Execute scripts conditionally",
		},
		{
			Label:         "incr  ",
			Detail:        "incr varName ?increment?",
			Documentation: " Increment the value of a variable",
		},
		{
			Label:         "join ",
			Detail:        "join list ?joinString?",
			Documentation: "Create a string by joining together list elements",
		},
		{
			Label:         "lappend ",
			Detail:        "lappend varName ?value value value ...?",
			Documentation: "Append list elements onto a variable",
		},
		{
			Label:         "lassign ",
			Detail:        "lassign list ?varName ...?",
			Documentation: "Assign list elements to variables",
		},
		{
			Label:         "linsert  ",
			Detail:        "linsert list index ?element element ...?",
			Documentation: "Insert elements into a list",
		},
		{
			Label:         "list ",
			Detail:        "list ?arg arg ...?",
			Documentation: " Create a list",
		},
		{
			Label:         "my",
			Detail:        "my methodName ?arg ...?",
			Documentation: " invoke any method of current object or its class",
		},
		{
			Label:         "proc ",
			Detail:        "proc name args body",
			Documentation: "Create a Tcl procedure",
		},
		{
			Label:         "read ",
			Detail:        "read ?-nonewline? channelId",
			Documentation: "Read from a channel",
		},
		{
			Label:         "rename",
			Detail:        "rename oldName newName",
			Documentation: "Rename or delete a command",
		},
		{
			Label:         "return",
			Detail:        "return",
			Documentation: "Return from a procedure, or set return code of a script",
		},
		{
			Label:         "set ",
			Detail:        "set varName ?value?",
			Documentation: "Read and write variables",
		},
		{
			Label:         "time",
			Detail:        "time script ?count?",
			Documentation: " Time the execution of a script",
		},
		{
			Label:         "try",
			Detail:        "try body ?handler...? ?finally script?",
			Documentation: " Trap and process errors and exceptions",
		},
		{
			Label:         "flush",
			Detail:        "flush channelId",
			Documentation: " The flush command forces any buffered output for the specified channel to be written immediately. ",
		},
		{
			Label:         "else",
			Detail:        "used in loops",
			Documentation: "",
		},
		{
			Label:         "string",
			Detail:        "",
			Documentation: "",
		},
		{
			Label:         "list",
			Detail:        "",
			Documentation: "",
		},
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
