package qtwidgets

// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDoubleSpinBox struct {
	*QAbstractSpinBox
}

func (this *QDoubleSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QDoubleSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQDoubleSpinBoxFromPointer(cthis unsafe.Pointer) *QDoubleSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QDoubleSpinBox{bcthis0}
}
func (*QDoubleSpinBox) NewFromPointer(cthis unsafe.Pointer) *QDoubleSpinBox {
	return NewQDoubleSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:115
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDoubleSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:126
// index:0
// Public
// void QDoubleSpinBox(class QWidget *)
func NewQDoubleSpinBox(parent *QWidget /*777 QWidget **/) *QDoubleSpinBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDoubleSpinBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:127
// index:0
// Public virtual
// void ~QDoubleSpinBox()
func DeleteQDoubleSpinBox(*QDoubleSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:129
// index:0
// Public
// double value()
func (this *QDoubleSpinBox) Value() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:131
// index:0
// Public
// QString prefix()
func (this *QDoubleSpinBox) Prefix() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6prefixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:132
// index:0
// Public
// void setPrefix(const class QString &)
func (this *QDoubleSpinBox) SetPrefix(prefix *qtcore.QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setPrefixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:134
// index:0
// Public
// QString suffix()
func (this *QDoubleSpinBox) Suffix() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6suffixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:135
// index:0
// Public
// void setSuffix(const class QString &)
func (this *QDoubleSpinBox) SetSuffix(suffix *qtcore.QString) {
	var convArg0 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setSuffixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:137
// index:0
// Public
// QString cleanText()
func (this *QDoubleSpinBox) CleanText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox9cleanTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:139
// index:0
// Public
// double singleStep()
func (this *QDoubleSpinBox) SingleStep() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10singleStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:140
// index:0
// Public
// void setSingleStep(double)
func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox13setSingleStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:142
// index:0
// Public
// double minimum()
func (this *QDoubleSpinBox) Minimum() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:143
// index:0
// Public
// void setMinimum(double)
func (this *QDoubleSpinBox) SetMinimum(min float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMinimumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:145
// index:0
// Public
// double maximum()
func (this *QDoubleSpinBox) Maximum() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:146
// index:0
// Public
// void setMaximum(double)
func (this *QDoubleSpinBox) SetMaximum(max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMaximumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:148
// index:0
// Public
// void setRange(double, double)
func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setRangeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:150
// index:0
// Public
// int decimals()
func (this *QDoubleSpinBox) Decimals() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8decimalsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:151
// index:0
// Public
// void setDecimals(int)
func (this *QDoubleSpinBox) SetDecimals(prec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox11setDecimalsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:153
// index:0
// Public virtual
// QValidator::State validate(class QString &, int &)
func (this *QDoubleSpinBox) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:154
// index:0
// Public virtual
// double valueFromText(const class QString &)
func (this *QDoubleSpinBox) ValueFromText(text *qtcore.QString) float64 {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13valueFromTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qspinbox.h:155
// index:0
// Public virtual
// QString textFromValue(double)
func (this *QDoubleSpinBox) TextFromValue(val float64) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13textFromValueEd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:156
// index:0
// Public virtual
// void fixup(class QString &)
func (this *QDoubleSpinBox) Fixup(str *qtcore.QString) {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:159
// index:0
// Public
// void setValue(double)
func (this *QDoubleSpinBox) SetValue(val float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setValueEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:162
// index:0
// Public
// void valueChanged(double)
func (this *QDoubleSpinBox) ValueChanged(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:163
// index:1
// Public
// void valueChanged(const class QString &)
func (this *QDoubleSpinBox) ValueChanged_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end