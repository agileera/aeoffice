package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const AEPROJECT = ".aeproject"
const AEDOCSUFFIX = ".aedoc"
const OUTSUFFIX = ".docx"
const TARGETPATHBASE = "dist"
const TEMPLATEHBASE = "template"

func printUsage(args []string) {
	log.Println("Usage:")
	log.Println(args[0] + " [some.aedoc]|[all]")
}

func getProjectRoot(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	stat, statErr := os.Stat(absPath)
	if os.IsNotExist(statErr) {
		log.Println(err)
		os.Exit(-1)
	}

	if stat.IsDir() {
		return _getProjectRoot(absPath)
	} else {
		parentDir, _ := filepath.Split(absPath)
		return _getProjectRoot(parentDir)
	}

}

func _getProjectRoot(absPath string) (string, error) {
	_, statErr := os.Stat(filepath.Join(absPath, AEPROJECT))
	if os.IsNotExist(statErr) {
		dir := filepath.Dir(absPath)
		if dir == absPath {
			return "", nil
		} else {
			return _getProjectRoot(dir)
		}
	} else {
		return absPath, nil
	}
}

func aebuildFile(path, targetPath, template string) error {
	log.Println("AEBuild: " + path)
	basepath := filepath.Dir(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	includeFilePaths := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			hasBrackets := strings.Contains(line, "(") && strings.Contains(line, ")")
			if hasBrackets {
				filePath := line[strings.Index(line, "(")+1 : strings.Index(line, ")")]
				relativeFilePath := filepath.Join(basepath, filePath)
				includeFilePaths = append(includeFilePaths, relativeFilePath)
				log.Println("Include: " + relativeFilePath)
			}
		}
	}
	_aebuildFile(includeFilePaths, targetPath, template)

	return nil
}

func _aebuildFile(includeFilePaths []string, targetPath string, template string) {
	args := []string{"-o", targetPath}
	if template != "" {
		args = append(args, "--reference-doc")
		args = append(args, template)
	}
	args = append(args, includeFilePaths...)
	cmd := exec.Command("pandoc", args...)
	log.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

}

func mkTargetPath(projectRoot, path string) string {
	absPath, _ := filepath.Abs(path)
	relativePath, _ := filepath.Rel(projectRoot, absPath)
	targetPath := filepath.Join(projectRoot, TARGETPATHBASE,
		strings.Replace(relativePath, AEDOCSUFFIX, OUTSUFFIX, -1))
	err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	// log.Println(targetPath)
	return targetPath
}

func getTemplatePath(projectRoot string) string {
	templatePath := filepath.Join(projectRoot, TEMPLATEHBASE, TEMPLATEHBASE+OUTSUFFIX)
	_, err := os.Stat(templatePath)
	// log.Println(templatePath)
	if os.IsNotExist(err) {
		return ""
	} else {
		return templatePath
	}
}

func aebuild(path string) {
	projectRoot, err := getProjectRoot(path)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println("AEProject Root:", projectRoot)

	targetPath := mkTargetPath(projectRoot, path)
	template := getTemplatePath(projectRoot)
	log.Println("Target Path:", targetPath)
	log.Println("Template:", template)

	stat, statErr := os.Stat(path)
	if statErr != nil {
		log.Println(statErr)
		os.Exit(-1)
	}

	if !stat.IsDir() {
		if strings.HasSuffix(path, AEDOCSUFFIX) {

			aebuildFile(path, targetPath, template)
		} else {
			log.Println("Please provide a " + AEDOCSUFFIX + " file")
			os.Exit(-1)
		}
	}

}

func main() {
	args := os.Args
	// log.Println(len(args))
	if len(args) != 2 {
		printUsage(args)
		os.Exit(-1)
	}

	path := args[1]
	aebuild(path)

}
