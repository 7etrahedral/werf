	"os"
	unrecognized   parserState = "unrecognized"
	diffBegin      parserState = "diffBegin"
	diffBody       parserState = "diffBody"
	newFileDiff    parserState = "newFileDiff"
	deleteFileDiff parserState = "deleteFileDiff"
	modifyFileDiff parserState = "modifyFileDiff"
	ignoreDiff     parserState = "ignoreDiff"
func debugPatchParser() bool {
	return os.Getenv("DAPP_TRUE_GIT_DEBUG_PATCH_PARSER") == "1"
}

	if debugPatchParser() {
		oldState := p.state
		fmt.Printf("TRUE_GIT parse diff line: state=%#v line=%#v\n", oldState, line)
		defer func() {
			fmt.Printf("TRUE_GIT parse diff line: state change: %#v => %#v\n", oldState, p.state)
		}()
	}

			return p.handleModifyFileDiff(line)
		if strings.HasPrefix(line, "new mode ") {
			return p.writeOutLine(line)
		}
		if strings.HasPrefix(line, "diff --git ") {
			return p.handleDiffBegin(line)
		if strings.HasPrefix(line, "Submodule ") {
			return p.handleSubmoduleLine(line)
		}
		return p.writeOutLine(line)