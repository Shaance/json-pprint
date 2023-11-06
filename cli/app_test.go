package cli

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func Test_writeOutput(t *testing.T) {
	t.Run("writes to correct file given an outputfilepath", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tmpFile, err := os.CreateTemp("", "tmp")
		require.NoError(t, err)

		tmpFileName := tmpFile.Name()
		defer tmpFile.Close()
		defer os.Remove(tmpFileName)

		fileContent := "{'key' : 'value'}"
		expected := fileContent + "\n"

		inputFilePath := ""
		outputFilePath := tmpFileName
		err = writeOutput(fileContent, true, inputFilePath, tmpFileName, ActualOS{})
		require.NoError(t, err)

		actual, err := os.ReadFile(outputFilePath)

		require.NoError(t, err)
		require.Equal(t, expected, string(actual))
	})

	t.Run("writes to correct file given empty outputfilepath", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tmpFile, err := os.CreateTemp("", "tmp")
		require.NoError(t, err)

		tmpFileName := tmpFile.Name()
		defer tmpFile.Close()
		defer os.Remove(tmpFileName)

		fileContent := "{'key' : 'value'}"
		expected := fileContent + "\n"

		inputFilePath := tmpFileName
		outputFilePath := ""
		err = writeOutput(fileContent, true, inputFilePath, outputFilePath, ActualOS{})
		require.NoError(t, err)

		actual, err := os.ReadFile(inputFilePath)

		require.NoError(t, err)
		require.Equal(t, expected, string(actual))
	})

	t.Run("does not write to file if writeToFile is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fileContent := "{'key' : 'value'}"
		expected := true

		inputFilePath := ""
		outputFilePath := ""
		writeToFile := false

		err := writeOutput(fileContent, writeToFile, inputFilePath, outputFilePath, ActualOS{})
		require.NoError(t, err)

		_, err = os.ReadFile(outputFilePath)
		require.Equal(t, expected, os.IsNotExist(err))
	})
}

func Test_retrieveJsonInput(t *testing.T) {
	t.Run("should return first argument when filepath is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedOs := NewMockOSInterface(ctrl)
		firstArg := "{'key' : 'value'}"

		expected := firstArg
		actual, err := retrieveJsonInput(firstArg, "", mockedOs)

		require.NoError(t, err)
		require.Equal(t, actual, expected)
	})

	t.Run("should return file content when filepath is not empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedOs := NewMockOSInterface(ctrl)
		filePath := "/tmp/something/cool.json"
		fileContent := "{'key' : 'value'}"
		firstArg := ""

		mockedOs.
			EXPECT().
			ReadFile(gomock.Eq(filePath)).
			Return([]byte(fileContent), nil)

		expected := fileContent
		actual, err := retrieveJsonInput(firstArg, filePath, mockedOs)

		require.NoError(t, err)
		require.Equal(t, actual, expected)
	})

	t.Run("should error if reading file content fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedOs := NewMockOSInterface(ctrl)
		filePath := "/tmp/something/cool.json"
		firstArg := ""

		mockedOs.
			EXPECT().
			ReadFile(gomock.Eq(filePath)).
			Return(nil, errors.New("something bad"))

		actual, err := retrieveJsonInput(firstArg, filePath, mockedOs)

		require.Error(t, err)
		require.Equal(t, actual, "")
	})
}

func Test_indentJson(t *testing.T) {
	t.Run("should indent json using tab when useSpaces is false", func(t *testing.T) {
		rawJson := `{"a":"b"}`
		expected := `{
	"a": "b"
}`
		useSpaces := false

		actual, err := indentJson(rawJson, useSpaces)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("should indent json using 2 spaces when useSpaces is true", func(t *testing.T) {
		rawJson := `{"a":"b"}`
		expected := `{
  "a": "b"
}`
		useSpaces := true

		actual, err := indentJson(rawJson, useSpaces)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
