package pprint

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func Test_writeOutput(t *testing.T) {
	// t.Run("should write the buffer content to stdout when writeToFile is false", func(t *testing.T) {
	// 	writer := &bytes.Buffer{}
	// 	require.NoError(t, writeOutput(bytes.NewBufferString("test"), false, "./a.json", "./a.json", writer))
	// 	require.Equal(t, "test\n", writer.String())
	// })

	// Another test case where writing to a file is simulated could be written here
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
		buf := &bytes.Buffer{}
		useSpaces := false

		require.NoError(t, indentJson(rawJson, useSpaces, buf))
		require.Equal(t, expected, buf.String())
	})

	t.Run("should indent json using 2 spaces when useSpaces is true", func(t *testing.T) {
		rawJson := `{"a":"b"}`
		expected := `{
  "a": "b"
}`
		buf := &bytes.Buffer{}
		useSpaces := true

		require.NoError(t, indentJson(rawJson, useSpaces, buf))
		require.Equal(t, expected, buf.String())
	})
}
