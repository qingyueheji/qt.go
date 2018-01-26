package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QObject struct {
	*qtrt.CObject
}

func (this *QObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObject) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQObjectFromPointer(cthis unsafe.Pointer) *QObject {
	return &QObject{&qtrt.CObject{cthis}}
}
func (*QObject) NewFromPointer(cthis unsafe.Pointer) *QObject {
	return NewQObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:118
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QObject) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:123
// index:0
// Public
// void QObject(class QObject *)
func NewQObject(parent *QObject /*777 QObject **/) *QObject {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObjectC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQObjectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// Public virtual
// void ~QObject()
func DeleteQObject(*QObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObjectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:126
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QObject) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:127
// index:0
// Public virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QObject) EventFilter(watched *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11eventFilterEPS_P6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:144
// index:0
// Public
// QString objectName()
func (this *QObject) ObjectName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10objectNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:145
// index:0
// Public
// void setObjectName(const class QString &)
func (this *QObject) SetObjectName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject13setObjectNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:147
// index:0
// Public inline
// bool isWidgetType()
func (this *QObject) IsWidgetType() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject12isWidgetTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:148
// index:0
// Public inline
// bool isWindowType()
func (this *QObject) IsWindowType() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject12isWindowTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:150
// index:0
// Public inline
// bool signalsBlocked()
func (this *QObject) SignalsBlocked() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14signalsBlockedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:151
// index:0
// Public
// bool blockSignals(_Bool)
func (this *QObject) BlockSignals(b bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject12blockSignalsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:153
// index:0
// Public
// QThread * thread()
func (this *QObject) Thread() *QThread /*777 QThread **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject6threadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQThreadFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:154
// index:0
// Public
// void moveToThread(class QThread *)
func (this *QObject) MoveToThread(thread *QThread /*777 QThread **/) {
	var convArg0 = thread.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject12moveToThreadEP7QThread", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:156
// index:0
// Public
// int startTimer(int, Qt::TimerType)
func (this *QObject) StartTimer(interval int, timerType int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobject.h:164
// index:0
// Public
// void killTimer(int)
func (this *QObject) KillTimer(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9killTimerEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:209
// index:0
// Public
// void setParent(class QObject *)
func (this *QObject) SetParent(parent *QObject /*777 QObject **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9setParentEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:210
// index:0
// Public
// void installEventFilter(class QObject *)
func (this *QObject) InstallEventFilter(filterObj *QObject /*777 QObject **/) {
	var convArg0 = filterObj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject18installEventFilterEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:211
// index:0
// Public
// void removeEventFilter(class QObject *)
func (this *QObject) RemoveEventFilter(obj *QObject /*777 QObject **/) {
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject17removeEventFilterEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:213
// index:0
// Public static
// QMetaObject::Connection connect(const class QObject *, const char *, const class QObject *, const char *, Qt::ConnectionType)
func (this *QObject) Connect(sender *QObject /*777 const QObject **/, signal string, receiver *QObject /*777 const QObject **/, member string, arg4 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_POINTER, sender, signal, receiver, member, arg4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QObject_Connect(sender *QObject /*777 const QObject **/, signal string, receiver *QObject /*777 const QObject **/, member string, arg4 int) int {
	var nilthis *QObject
	rv := nilthis.Connect(sender, signal, receiver, member, arg4)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:216
// index:1
// Public static
// QMetaObject::Connection connect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &, Qt::ConnectionType)
func (this *QObject) Connect_1(sender *QObject /*777 const QObject **/, signal *QMetaMethod, receiver *QObject /*777 const QObject **/, method *QMetaMethod, type_ int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_POINTER, sender, signal, receiver, method, type_)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QObject_Connect_1(sender *QObject /*777 const QObject **/, signal *QMetaMethod, receiver *QObject /*777 const QObject **/, method *QMetaMethod, type_ int) int {
	var nilthis *QObject
	rv := nilthis.Connect_1(sender, signal, receiver, method, type_)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:220
// index:2
// Public inline
// QMetaObject::Connection connect(const class QObject *, const char *, const char *, Qt::ConnectionType)
func (this *QObject) Connect_2(sender *QObject /*777 const QObject **/, signal string, member string, type_ int) int {
	var convArg0 = sender.GetCthis()
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, type_)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:342
// index:0
// Public static
// bool disconnect(const class QObject *, const char *, const class QObject *, const char *)
func (this *QObject) Disconnect(sender *QObject /*777 const QObject **/, signal string, receiver *QObject /*777 const QObject **/, member string) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_PKcS1_S3_", ffiqt.FFI_TYPE_POINTER, sender, signal, receiver, member)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QObject_Disconnect(sender *QObject /*777 const QObject **/, signal string, receiver *QObject /*777 const QObject **/, member string) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:344
// index:1
// Public static
// bool disconnect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &)
func (this *QObject) Disconnect_1(sender *QObject /*777 const QObject **/, signal *QMetaMethod, receiver *QObject /*777 const QObject **/, member *QMetaMethod) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_", ffiqt.FFI_TYPE_POINTER, sender, signal, receiver, member)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QObject_Disconnect_1(sender *QObject /*777 const QObject **/, signal *QMetaMethod, receiver *QObject /*777 const QObject **/, member *QMetaMethod) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect_1(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:346
// index:2
// Public inline
// bool disconnect(const char *, const class QObject *, const char *)
func (this *QObject) Disconnect_2(signal string, receiver *QObject /*777 const QObject **/, member string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:349
// index:3
// Public inline
// bool disconnect(const class QObject *, const char *)
func (this *QObject) Disconnect_3(receiver *QObject /*777 const QObject **/, member string) bool {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:390
// index:0
// Public
// void dumpObjectTree()
func (this *QObject) DumpObjectTree() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject14dumpObjectTreeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:393
// index:1
// Public
// void dumpObjectTree()
func (this *QObject) DumpObjectTree_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14dumpObjectTreeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:391
// index:0
// Public
// void dumpObjectInfo()
func (this *QObject) DumpObjectInfo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject14dumpObjectInfoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:394
// index:1
// Public
// void dumpObjectInfo()
func (this *QObject) DumpObjectInfo_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14dumpObjectInfoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:397
// index:0
// Public
// bool setProperty(const char *, const class QVariant &)
func (this *QObject) SetProperty(name string, value *QVariant) bool {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11setPropertyEPKcRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:398
// index:0
// Public
// QVariant property(const char *)
func (this *QObject) Property(name string) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8propertyEPKc", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:403
// index:0
// Public static
// uint registerUserData()
func (this *QObject) RegisterUserData() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject16registerUserDataEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QObject_RegisterUserData() uint {
	var nilthis *QObject
	rv := nilthis.RegisterUserData()
	return rv
}

// /usr/include/qt/QtCore/qobject.h:404
// index:0
// Public
// void setUserData(uint, class QObjectUserData *)
func (this *QObject) SetUserData(id uint, data *QObjectUserData /*777 QObjectUserData **/) {
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11setUserDataEjP15QObjectUserData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:405
// index:0
// Public
// QObjectUserData * userData(uint)
func (this *QObject) UserData(id uint) *QObjectUserData /*777 QObjectUserData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8userDataEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:409
// index:0
// Public
// void destroyed(class QObject *)
func (this *QObject) Destroyed(arg0 *QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:413
// index:0
// Public inline
// QObject * parent()
func (this *QObject) Parent() *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:415
// index:0
// Public inline
// bool inherits(const char *)
func (this *QObject) Inherits(classname string) bool {
	var convArg0 = qtrt.CString(classname)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8inheritsEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:419
// index:0
// Public
// void deleteLater()
func (this *QObject) DeleteLater() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11deleteLaterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:422
// index:0
// Protected
// QObject * sender()
func (this *QObject) Sender() *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject6senderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:423
// index:0
// Protected
// int senderSignalIndex()
func (this *QObject) SenderSignalIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject17senderSignalIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobject.h:424
// index:0
// Protected
// int receivers(const char *)
func (this *QObject) Receivers(signal string) int {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject9receiversEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobject.h:425
// index:0
// Protected
// bool isSignalConnected(const class QMetaMethod &)
func (this *QObject) IsSignalConnected(signal *QMetaMethod) bool {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject17isSignalConnectedERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:427
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QObject) TimerEvent(event *QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:428
// index:0
// Protected virtual
// void childEvent(class QChildEvent *)
func (this *QObject) ChildEvent(event *QChildEvent /*777 QChildEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:429
// index:0
// Protected virtual
// void customEvent(class QEvent *)
func (this *QObject) CustomEvent(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11customEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:431
// index:0
// Protected virtual
// void connectNotify(const class QMetaMethod &)
func (this *QObject) ConnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:432
// index:0
// Protected virtual
// void disconnectNotify(const class QMetaMethod &)
func (this *QObject) DisconnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end