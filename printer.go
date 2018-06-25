package j2h

import (
	"fmt"
	"strings"
)

// Printer is an interface that prints ddl of hive.
type Printer interface {
	Print() string
}

// PrimitivePrinter is a printer structure corresponding to the primitive type of hive.
type PrimitivePrinter struct {
	depth     int
	colName   string
	typeName  string
	delimiter string
}

// StructPrinter is a printer structure corresponding to the struct type of hive.
type StructPrinter struct {
	depth     int
	colName   string
	typeName  string
	delimiter string
	members   []Printer
}

// ArrayPrinter is a printer structure corresponding to the array type of hive.
type ArrayPrinter struct {
	headerDepth   int
	fotterDepth   int
	colName       string
	typeName      string
	delimiter     string
	headerNewline string
	fotterNewline string
	member        Printer
}

// NewPrimitivePrinter creates and returns a new PrimitivePrinter.
func NewPrimitivePrinter(depth int, colName, typeName, delimiter string) *PrimitivePrinter {
	return &PrimitivePrinter{
		depth:     depth,
		colName:   colName,
		typeName:  typeName,
		delimiter: delimiter,
	}
}

// NewStructPrinter creates and returns a new StructPrinter.
func NewStructPrinter(depth int, colName, delimiter string, plist []Printer) *StructPrinter {
	return &StructPrinter{
		depth:     depth,
		colName:   colName,
		typeName:  "struct",
		delimiter: delimiter,
		members:   plist,
	}
}

// NewPrimitiveArrayPrinter creates and returns a new ArrayPrinter whose element is a primitive type.
func NewPrimitiveArrayPrinter(depth int, colName, delimiter, elementType string) *ArrayPrinter {
	p := NewPrimitivePrinter(0, "", elementType, "")

	return &ArrayPrinter{
		headerDepth:   depth,
		fotterDepth:   0,
		colName:       colName,
		typeName:      "array",
		delimiter:     delimiter,
		headerNewline: "",
		fotterNewline: "",
		member:        p,
	}
}

// NewStructArrayPrinter creates and returns a new ArrayPrinter whose element is a struct type.
func NewStructArrayPrinter(depth int, colName, delimiter string, plist []Printer) *ArrayPrinter {
	p := NewStructPrinter(depth+1, "", "", plist)

	return &ArrayPrinter{
		headerDepth:   depth,
		fotterDepth:   depth,
		colName:       colName,
		typeName:      "array",
		delimiter:     delimiter,
		headerNewline: "\n",
		fotterNewline: "\n",
		member:        p,
	}
}

// NewMultipleArrayPrinter creates and returns a new ArrayPrinter whose element is a array type.
func NewMultipleArrayPrinter(depth int, colName, delimiter string, p Printer) *ArrayPrinter {
	return &ArrayPrinter{
		headerDepth:   depth,
		fotterDepth:   depth,
		colName:       colName,
		typeName:      "array",
		delimiter:     delimiter,
		headerNewline: "\n",
		fotterNewline: "\n",
		member:        p,
	}
}

// PrintHeader prints header of hive ddl.
func PrintHeader() string {
	return fmt.Sprint("(")
}

// PrintFooter  prints footer of hive ddl.
func PrintFooter() string {
	return fmt.Sprint(")")
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p PrimitivePrinter) Print() string {
	return fmt.Sprintf("%s%s%s", p.colName, p.delimiter, p.typeName)
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p StructPrinter) Print() string {
	structPirntHeader := fmt.Sprintf("%s%s%s<", p.colName, p.delimiter, p.typeName)
	structPirntFooter := ">"

	var mPrints []string
	for _, v := range p.members {
		mPrints = append(mPrints, v.Print())
	}
	mPrint := strings.Join(mPrints, ",")

	return structPirntHeader + mPrint + structPirntFooter
}

// Print prints one line of hive ddl corresponding to the array type.
func (p ArrayPrinter) Print() string {
	structPirntHeader := fmt.Sprintf("%s%s%s<", p.colName, p.delimiter, p.typeName)
	structPirntFooter := ">"

	return structPirntHeader + p.member.Print() + structPirntFooter
}

func printIndent(depth int) string {
	return strings.Repeat("  ", depth)
}
