package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.IsIPAddress

	// Comment ignored

	//lintignore:V004
	validation.SingleIP()

	validation.SingleIP() //lintignore:V004

	// Failing

	validation.SingleIP() // want "deprecated validation.SingleIP should be replaced with validation.IsIPAddress"
}
