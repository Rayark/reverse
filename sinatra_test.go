package reverse

import (
	"testing"
)

func TestExtracSinatraParams(t *testing.T) {

	params := extractSinatraParams("/:page/:line/edit")

	if params[0] != ":page" {
		t.Fatal(":page cannot extracted")
	}

	if params[1] != ":line" {
		t.Fatal(":line cannot extracted")
	}
}

func TestSinatraAdd(t *testing.T) {

	root := NewURLStore()
	root.S("APIIndex", "/")

	admin := NewURLStore()
	admin.S("AdminIndex", "/")
	admin.S("AdminPost", "/:post")
	admin.S("AdminEditPost", "/:post/edit")

	err := root.Append("/admin", admin)
	if err != nil {
		panic(err)
	}

	if root.Rev("APIIndex") != "/" {
		t.Fatal("APIIndex")
	}

	if root.Rev("AdminIndex") != "/admin/" {
		t.Fatal("AdminIndex")
	}

	if root.Rev("AdminPost", "test") != "/admin/test" {
		t.Fatal("AdminPost")
	}

	if root.Rev("AdminEditPost", "test") != "/admin/test/edit" {
		t.Fatal("AdminEditPost")
	}

}
