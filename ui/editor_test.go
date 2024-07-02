package ui

import "testing"

var hell = Editor{
	text: []editorLine{{
		runes:    []rune{'h', 'e', 'l', 'l'},
		clusters: []int{0, 1, 2, 3, 4},
	}},
	textWidth: []int{0, 1, 2, 3, 4},
	cursorIdx: 4,
	offsetIdx: 0,
	width:     5,
}

func assertEditorEq(t *testing.T, actual, expected Editor) {
	if len(actual.text) != len(expected.text) {
		t.Errorf("expected text len to be %d, got %d\n", len(expected.text), len(actual.text))
	} else {
		for i := 0; i < len(actual.text[0].clusters); i++ {
			a := actual.text[0].clusters[i]
			e := expected.text[0].clusters[i]

			if a != e {
				t.Errorf("expected cluster #%d to be '%c', got '%c'\n", i, e, a)
			}
		}
		for i := 0; i < len(actual.text[0].runes); i++ {
			a := actual.text[0].runes[i]
			e := expected.text[0].runes[i]

			if a != e {
				t.Errorf("expected rune #%d to be '%c', got '%c'\n", i, e, a)
			}
		}
	}

	if len(actual.textWidth) != len(expected.textWidth) {
		t.Errorf("expected textWidth len to be %d, got %d\n", len(expected.textWidth), len(actual.textWidth))
	} else {
		for i := 0; i < len(actual.textWidth); i++ {
			a := actual.textWidth[i]
			e := expected.textWidth[i]

			if a != e {
				t.Errorf("expected width #%d to be %d, got %d\n", i, e, a)
			}
		}
	}

	if actual.cursorIdx != expected.cursorIdx {
		t.Errorf("expected cursorIdx to be %d, got %d\n", expected.cursorIdx, actual.cursorIdx)
	}

	if actual.offsetIdx != expected.offsetIdx {
		t.Errorf("expected offsetIdx to be %d, got %d\n", expected.offsetIdx, actual.offsetIdx)
	}

	if actual.width != expected.width {
		t.Errorf("expected width to be %d, got %d\n", expected.width, actual.width)
	}
}

func TestOneLetter(t *testing.T) {
	e := NewEditor(&UI{})
	e.Resize(5)
	e.PutRune('h')
	assertEditorEq(t, e, Editor{
		text: []editorLine{{
			runes:    []rune{'h'},
			clusters: []int{0, 1},
		}},
		textWidth: []int{0, 1},
		cursorIdx: 1,
		offsetIdx: 0,
		width:     5,
	})
}

func TestFourLetters(t *testing.T) {
	e := NewEditor(&UI{})
	e.Resize(5)
	e.PutRune('h')
	e.PutRune('e')
	e.PutRune('l')
	e.PutRune('l')
	assertEditorEq(t, e, hell)
}

func TestOneLeft(t *testing.T) {
	e := NewEditor(&UI{})
	e.Resize(5)
	e.PutRune('h')
	e.PutRune('l')
	e.Left()
	e.PutRune('e')
	e.PutRune('l')
	e.Right()
	assertEditorEq(t, e, hell)
}

func TestOneRem(t *testing.T) {
	e := NewEditor(&UI{})
	e.Resize(5)
	e.PutRune('h')
	e.PutRune('l')
	e.RemCluster()
	e.PutRune('e')
	e.PutRune('l')
	e.PutRune('l')
	assertEditorEq(t, e, hell)
}

func TestLeftAndRem(t *testing.T) {
	e := NewEditor(&UI{})
	e.Resize(5)
	e.PutRune('h')
	e.PutRune('l')
	e.PutRune('e')
	e.Left()
	e.RemCluster()
	e.Right()
	e.PutRune('l')
	e.PutRune('l')
	assertEditorEq(t, e, hell)
}
