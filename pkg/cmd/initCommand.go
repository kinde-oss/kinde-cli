package cmd

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Template struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	Branch      string `json:"branch"`
	Path        string `json:"path"`
}

type KindeConfig struct {
	RootDir string `json:"rootDir"`
	Version string `json:"version"`
}

var templates = map[string]Template{
	"orbit": {
		Name:        "Orbit Template",
		Description: "A React starter template",
		Repo:        "kinde-starter-kits/custom-ui-orbit",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"splitscape": {
		Name:        "Splitscape Template",
		Description: "A React starter template",
		Repo:        "kinde-starter-kits/custom-ui-splitscape",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"evolve-ai": {
		Name:        "Evolve.ai Template",
		Description: "A React starter template",
		Repo:        "kinde-starter-kits/custom-ui-evolve-ai",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"bark-n-bite": {
		Name:        "Bark & Bite Template",
		Description: "A React starter template",
		Repo:        "kinde-starter-kits/custom-ui-barknbite",
		Branch:      "main",
		Path:        "kindeSrc",
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Kinde projects",
	Long:  "Initialize various types of Kinde projects",
}

// customUICmd represents the custom-ui subcommand
type customUICmd struct {
	cmd      *cobra.Command
	rootDir  string
	template string
}

func newCustomUICmd() *customUICmd {
	c := &customUICmd{}
	cmd := &cobra.Command{
		Use:   "custom-ui",
		Short: "Initialize a new project with Kinde Custom UI",
		Long:  "Initialize a new project with Kinde Custom UI templates",
		RunE:  c.run,
	}

	cmd.Flags().StringVar(&c.rootDir, "rootDir", "kindeSrc", "Specify the root directory name")
  cmd.Flags().StringVar(&c.template, "template", "", "Specify the template (orbit, splitscape, evolve-ai, or bark-n-bite)")

	c.cmd = cmd
	return c
}

func (c *customUICmd) run(cmd *cobra.Command, args []string) error {
	fmt.Println("\n🚀 Getting started with Kinde Custom UIs")

	shouldProceed, err := c.checkExistingSetup()
	if err != nil {
		return err
	}

	if !shouldProceed {
		fmt.Println("\nSetup cancelled. Your existing Kinde configuration remains unchanged.")
		return nil
	}

	// If flags weren't provided, prompt for input
	if c.rootDir == "kindeSrc" {
		validate := func(input string) error {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9-_]+$", input)
			if !matched {
				return fmt.Errorf("please use only letters, numbers, hyphens, and underscores")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "What is the root directory for your kinde code?",
			Default:  "kindeSrc",
			Validate: validate,
		}

		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %w", err)
		}
		c.rootDir = result
	}

	if c.template == "" {
		type templateItem struct {
			Key         string
			DisplayName string
		}
		
		var items []templateItem
		var displayItems []string
		
		for key, tmpl := range templates {
			displayName := fmt.Sprintf("%s - %s", tmpl.Name, tmpl.Description)
			items = append(items, templateItem{Key: key, DisplayName: displayName})
			displayItems = append(displayItems, displayName)
		}
		
		prompt := promptui.Select{
			Label: "Select a template",
			Items: displayItems,
		}
		
		index, _, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %w", err)
		}
		
		c.template = items[index].Key
	}

	targetDir := filepath.Join(".", c.rootDir)
	
	if err := c.createFromGitTemplate(c.template, targetDir); err != nil {
		return fmt.Errorf("failed to create template: %w", err)
	}

	if err := c.createKindeConfig(); err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}


	fmt.Println("\n✨ Custom UI template created successfully!")
	fmt.Println("\nNext steps:")
	fmt.Printf("1. cd %s\n", c.rootDir)
	fmt.Println("2. Edit page.tsx to customise your authentication UI")
	fmt.Println("3. Sync your changes on GitHub and Kinde")

	return nil
}

func (c *customUICmd) checkExistingSetup() (bool, error) {
	if _, err := os.Stat("kinde.json"); err == nil {
		// File exists
		prompt := promptui.Prompt{
			Label:     "A Kinde setup already exists in this directory. Would you like to delete it and start fresh?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			return false, fmt.Errorf("prompt failed: %w", err)
		}

		if result == "y" {
			// Remove existing directory and config
			configData, err := os.ReadFile("kinde.json")
			if err != nil {
				return false, fmt.Errorf("failed to read config: %w", err)
			}

			var config KindeConfig
			if err := json.Unmarshal(configData, &config); err != nil {
				return false, fmt.Errorf("failed to parse config: %w", err)
			}

			oldDirPath := filepath.Join(".", config.RootDir)
			if err := os.RemoveAll(oldDirPath); err != nil {
				return false, fmt.Errorf("failed to remove old directory: %w", err)
			}

			if err := os.Remove("kinde.json"); err != nil {
				return false, fmt.Errorf("failed to remove old config: %w", err)
			}

			return true, nil
		}
		return false, nil
	}
	return true, nil
}

// LimitedWriter wraps an io.Writer with a byte limit
type LimitedWriter struct {
	Writer   io.Writer
	MaxBytes int64
	Written  int64
}

// Write implements the io.Writer interface
func (lw *LimitedWriter) Write(p []byte) (n int, err error) {
	if lw.Written+int64(len(p)) > lw.MaxBytes {
		return 0, fmt.Errorf("file too large: limit is %d bytes", lw.MaxBytes)
	}
	n, err = lw.Writer.Write(p)
	lw.Written += int64(n)
	return n, err
}

// unzip extracts a zip archive to a destination directory
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	// Get the root directory name inside the ZIP
	var rootDir string
	if len(r.File) > 0 {
		rootDir = filepath.Dir(r.File[0].Name)
		// If root directory has multiple levels, get only the first directory
		if rootDir != "." {
			parts := strings.Split(rootDir, string(os.PathSeparator))
			rootDir = parts[0]
		}
	}

	// Extract files
	for _, f := range r.File {
		// Skip directories
		if f.FileInfo().IsDir() {
			continue
		}

		// Calculate relative path (removing root directory)
		relPath := f.Name
		if rootDir != "." && strings.HasPrefix(relPath, rootDir) {
			relPath = strings.TrimPrefix(relPath, rootDir+string(os.PathSeparator))
		}

		// Create target file path
		targetPath := filepath.Join(dest, relPath)

		// Create directory for file if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		// Open source file
		rc, err := f.Open()
		if err != nil {
			return err
		}

		// Create target file
		targetFile, err := os.Create(targetPath)
		if err != nil {
			rc.Close()
			return err
		}

		// Copy content
		_, err = io.Copy(targetFile, rc)
		targetFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}

	return nil
}


func (c *customUICmd) createFromGitTemplate(template, targetDir string) error {
	fmt.Println("Creating template from GitHub...")
	
	tmpl, ok := templates[template]
	if !ok {
		return fmt.Errorf("template %s not found", template)
	}

	// Create target directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create temporary directory for downloading
	tempDir, err := os.MkdirTemp("", "kinde-template-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Download zip file from GitHub API
	url := fmt.Sprintf("https://api.github.com/repos/%s/zipball/%s", tmpl.Repo, tmpl.Branch)
	fmt.Printf("Downloading template from %s...\n", url)
	
	// Create HTTP client
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set User-Agent header (GitHub API requires this)
	req.Header.Set("User-Agent", "Kinde-CLI")
	
	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download template: %w", err)
	}
	defer resp.Body.Close()
	
	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download template: HTTP status %s", resp.Status)
	}
	
	// Create zip file path
	zipPath := filepath.Join(tempDir, "template.zip")
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()
	
	// Limit download size to 50MB to prevent abuse
	limitedWriter := &LimitedWriter{
		Writer:   zipFile,
		MaxBytes: 50 * 1024 * 1024, // 50MB limit
	}
	
	// Download the zip file
	_, err = io.Copy(limitedWriter, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download zip file: %w", err)
	}
	
	// Extract the zip file
	extractPath := filepath.Join(tempDir, "extracted")
	if err := unzip(zipPath, extractPath); err != nil {
		return fmt.Errorf("failed to extract zip file: %w", err)
	}
	
	// Find the template path inside the extracted directory
	sourcePath := filepath.Join(extractPath, tmpl.Path)
	
	// Copy files from template to target directory
	if err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Calculate relative path
		relPath, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}
		
		// Skip root directory
		if relPath == "." {
			return nil
		}
		
		// Create target path
		destPath := filepath.Join(targetDir, relPath)
		
		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}
		
		// Copy file
		input, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read source file: %w", err)
		}
		
		if err := os.WriteFile(destPath, input, 0644); err != nil {
			return fmt.Errorf("failed to write destination file: %w", err)
		}
		
		return nil
	}); err != nil {
		return fmt.Errorf("failed to copy template files: %w", err)
	}
	
	fmt.Println("Template files copied successfully!")
	return nil
}


func (c *customUICmd) createKindeConfig() error {
	config := KindeConfig{
		RootDir: c.rootDir,
		Version: "2024-12-09",
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile("kinde.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}


func init() {
	// Add init command to root command
	rootCmd.AddCommand(initCmd)
	
	// Add custom-ui command to init command
	customUICmd := newCustomUICmd()
	initCmd.AddCommand(customUICmd.cmd)
}