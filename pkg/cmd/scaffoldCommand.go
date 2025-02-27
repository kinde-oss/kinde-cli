package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

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
		Description: "A simple starter template with essential Kinde authentication",
		Repo:        "kinde-starter-kits/custom-ui-orbit",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"splitscape": {
		Name:        "Splitscape Template",
		Description: "A simple starter template with essential Kinde authentication",
		Repo:        "kinde-starter-kits/custom-ui-splitscape",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"evolve-ai": {
		Name:        "Evolve.ai Template",
		Description: "A simple starter template with essential Kinde authentication",
		Repo:        "kinde-starter-kits/custom-ui-evolve-ai",
		Branch:      "main",
		Path:        "kindeSrc",
	},
	"bark-n-bite": {
		Name:        "Bark & Bite Template",
		Description: "A simple starter template with essential Kinde authentication",
		Repo:        "kinde-starter-kits/custom-ui-barknbite",
		Branch:      "main",
		Path:        "kindeSrc",
	},
}

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffold Kinde projects",
	Long:  "Scaffold various types of Kinde projects",
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

	if err := c.installDependencies(); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
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

	// Create temporary directory for cloning
	tempDir, err := os.MkdirTemp("", "kinde-template-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone the repository using git command
	repoURL := fmt.Sprintf("https://github.com/%s.git", tmpl.Repo)
	gitCmd := exec.Command("git", "clone", "--depth", "1", "--branch", tmpl.Branch, repoURL, tempDir)
	if output, err := gitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone repository: %s: %w", string(output), err)
	}

	// Copy files from template to target directory
	srcPath := filepath.Join(tempDir, tmpl.Path)
	if err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(srcPath, path)
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

func (c *customUICmd) installDependencies() error {
	fmt.Println("Installing dependencies...")

	// Define dependencies to install
	dependencies := []string{
		"@kinde/infrastructure",
		"react",
		"react-dom",
	}

	// Change to project directory
	if err := os.Chdir(c.rootDir); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}
	defer os.Chdir("..")

	// Initialize package.json if it doesn't exist
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		npmInitCmd := exec.Command("npm", "init", "-y")
		if output, err := npmInitCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to initialize package.json: %s: %w", string(output), err)
		}
	}

	// Install dependencies
	args := append([]string{"install", "--save"}, dependencies...)
	installCmd := exec.Command("npm", args...)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	fmt.Println("Dependencies installed successfully!")
	return nil
}

func init() {
	// Add scaffold command to root command
	rootCmd.AddCommand(scaffoldCmd)
	
	// Add custom-ui command to scaffold command
	customUICmd := newCustomUICmd()
	scaffoldCmd.AddCommand(customUICmd.cmd)
}