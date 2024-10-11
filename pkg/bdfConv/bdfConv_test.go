package bdfConv

import "testing"

func TestName(t *testing.T) {
	r := CreateTargetCFont("opt/wenquanyi_10pt.bdf", "32-123", "dextercai_test", "opt/export.c")
	t.Log(r)
}
