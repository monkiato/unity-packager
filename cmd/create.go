/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unity-packager/tools"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	assetsPath      string
	output          string
	addAssetsFolder bool
	ignoreFilters   []string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a .unitypackage file",
	Run: func(cmd *cobra.Command, args []string) {
		err := func() error {
			if !tools.FileExists(assetsPath) {
				return fmt.Errorf("Assets path doesn't exists")
			}

			tmpFolder := fmt.Sprintf("%s/%s/", cachedir, uuid.NewString())
			os.MkdirAll(tmpFolder, os.ModeDir|os.ModePerm)
			defer os.RemoveAll(tmpFolder)

			filepath.Walk(assetsPath, func(path string, info fs.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}

				//ignore meta files here as they will be evaluated later
				if filepath.Ext(path) == ".meta" {
					return nil
				}
				// ignore MacOS finder files
				if info.Name() == ".DS_Store" {
					return nil
				}

				for _, filter := range ignoreFilters {
					if strings.Contains(path, filter) {
						// filtered
						return nil
					}
				}
				// obtain file GUID
				guid, err := tools.GetGUID(path, true)
				if err != nil {
					return err
				}

				// create .meta file if not exists
				metadataPath := path + ".meta"
				if !tools.FileExists(metadataPath) {
					tools.CreateMetadata(path, guid)
				}

				// write 'asset' file
				if err := tools.CopyFile(path, fmt.Sprintf("%s/%s/asset", tmpFolder, guid), true); err != nil {
					fmt.Printf("Unable to copy file %s. %s", path, err)
					return err
				}
				fmt.Println("packaging file: " + path)

				// write 'asset.meta' file
				if err := tools.CopyFile(metadataPath, fmt.Sprintf("%s/%s/asset.meta", tmpFolder, guid), true); err != nil {
					fmt.Printf("Unable to copy metadata file %s. %s", metadataPath, err)
					return err
				}

				// write 'pathname' file
				relativePath, err := filepath.Rel(filepath.Dir(assetsPath), path)
				if err != nil {
					return err
				}
				if addAssetsFolder {
					relativePath = "Assets/" + relativePath
				}
				fmt.Println(relativePath)
				os.WriteFile(fmt.Sprintf("%s/%s/pathname", tmpFolder, guid), []byte(relativePath), os.ModePerm)

				return nil
			})

			return generatePackage(tmpFolder, output)
		}()

		if err != nil {
			fmt.Println("An error ocurred: " + err.Error())
		}
	},
}

func generatePackage(path string, filename string) error {
	unitypackageName := filename + ".unitypackage"
	// Create output file
	out, err := os.Create(unitypackageName)
	if err != nil {
		log.Fatalln("Error writing archive:", err)
		return err
	}
	defer out.Close()

	var files []string
	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalln("Error trying to list of files to package:", err)
		return err
	}

	// Create the archive and write the output to the "out" Writer
	err = tools.CreateArchive(files, out, path)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
		return err
	}

	fmt.Println("Unity package has been created successfully.")
	fmt.Println("Filename: " + unitypackageName)
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&assetsPath, "path", "p", "", "Specify folder to export into the .unitypackage file")
	createCmd.Flags().StringVarP(&output, "output", "o", "", "Specify output filename (without extension)")
	createCmd.Flags().BoolVar(&addAssetsFolder, "add-assets-folder", false, "Add Assets folder to the files included in the generated .unitypackage file")
	createCmd.Flags().StringSliceVarP(&ignoreFilters, "ignore", "i", []string{}, "Specify filename or extensions to exclude (e.g. --ingore \".csproj\")")

	createCmd.MarkFlagRequired("path")
	createCmd.MarkFlagRequired("output")
}
