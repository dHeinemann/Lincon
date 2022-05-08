package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRelative_RelativePath_ReturnsTrue(t *testing.T) {
	assert.True(t, IsRelative("foo.html"))
}

func TestIsRelative_AbsolutePath_ReturnsFalse(t *testing.T) {
	assert.False(t, IsRelative("/foo.html"))
}

func TestConvertToRelative_TargetInTopCurrentInDescendent(t *testing.T) {
	targetFile := "/b.html"
	currentFile := "/foo/bar/a.html"

	expected := "../../b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}

func TestConvertToRelative_BothInTop(t *testing.T) {
	targetFile := "/b.html"
	currentFile := "/a.html"

	expected := "b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}

func TestConvertToRelative_TargetInDescendentCurrentInTop(t *testing.T) {
	targetFile := "/foo/bar/b.html"
	currentFile := "/a.html"

	expected := "foo/bar/b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}

func TestConvertToRelative_BothInDifferentDescendents(t *testing.T) {
	targetFile := "/foo/bar/b.html"
	currentFile := "/foo/baz/a.html"

	expected := "../bar/b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}

func TestConvertToRelative_TargetInParent(t *testing.T) {
	targetFile := "/foo/b.html"
	currentFile := "/foo/bar/a.html"

	expected := "../b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}

func TestConvertToRelative_TargetInDescendent(t *testing.T) {
	targetFile := "/foo/bar/b.html"
	currentFile := "/foo/a.html"

	expected := "bar/b.html"
	assert.Equal(t, expected, ConvertToRelative(targetFile, currentFile, "/"))
}
