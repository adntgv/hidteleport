package main

import (
	"encoding/json"

	pb "github.com/adntgv/hidteleport/proto"
	hook "github.com/robotn/gohook"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func eventToBytesProto(evt hook.Event) ([]byte, error) {
	pbEvt := &pb.Event{
		Kind:     uint32(evt.Kind),
		When:     timestamppb.New(evt.When),
		Mask:     uint32(evt.Mask),
		Reserved: uint32(evt.Reserved),

		Keycode: uint32(evt.Keycode),
		Rawcode: uint32(evt.Rawcode),
		Keychar: evt.Keychar,

		Button: uint32(evt.Button),
		Clicks: uint32(evt.Clicks),

		X: int32(evt.X),
		Y: int32(evt.Y),

		Amount:    uint32(evt.Amount),
		Rotation:  evt.Rotation,
		Direction: uint32(evt.Direction),
	}
	return proto.Marshal(pbEvt)
}

func eventsFromBytesProto(bz []byte) (hook.Event, error) {
	pbEvt := &pb.Event{}
	err := proto.Unmarshal(bz, pbEvt)
	if err != nil {
		return hook.Event{}, nil
	}

	evt := hook.Event{
		Kind:     uint8(pbEvt.Kind),
		When:     pbEvt.When.AsTime(),
		Mask:     uint16(pbEvt.Mask),
		Reserved: uint16(pbEvt.Reserved),

		Keycode: uint16(pbEvt.Keycode),
		Rawcode: uint16(pbEvt.Rawcode),
		Keychar: pbEvt.Keychar,

		Button: uint16(pbEvt.Button),
		Clicks: uint16(pbEvt.Clicks),

		X: int16(pbEvt.X),
		Y: int16(pbEvt.Y),

		Amount:    uint16(pbEvt.Amount),
		Rotation:  pbEvt.Rotation,
		Direction: uint8(pbEvt.Direction),
	}

	return evt, nil
}

type screenSize struct {
	width  float32
	height float32
}

type scaledPosition struct {
	X float32
	Y float32
}

type delta struct {
	X int
	Y int
}

type eventWrapper struct {
	hook.Event
	ScaledMousePosition scaledPosition
	Delta               delta
}

func (ew eventWrapper) String() string {
	bz, _ := json.Marshal(ew)
	return ew.Event.String() + " " + string(bz)
}

var keysMap = map[uint16]string{
	//0:   "That key has no keycode",
	//3:   "break",
	8: "backspace",
	9: "tab",
	//12:  "clear",
	13: "enter",
	16: "shift",
	17: "ctrl",
	18: "alt",
	//19:  "pause/break",
	20: "capslock",
	//21:  "hangul",
	//25:  "hanja",
	27: "escape",
	//28:  "conversion",
	//29:  "non-conversion",
	32: "spacebar",
	33: "pageup",
	34: "pagedown",
	35: "end",
	36: "home",
	37: "left",
	38: "up",
	39: "right",
	40: "down",
	//41:  "select",
	//42:  "print",
	//43:  "execute",
	//44:  "Print Screen",
	45: "insert",
	46: "delete",
	//47:  "help",
	48: "num0",
	49: "num1",
	50: "num2",
	51: "num3",
	52: "num4",
	53: "num5",
	54: "num6",
	55: "num7",
	56: "num8",
	57: "num9",
	//58:  ":",
	//59:  "semicolon (firefox), equals",
	//60:  "<",
	61:  "equals (firefox)",
	63:  "ß",
	64:  "@ (firefox)",
	65:  "a",
	66:  "b",
	67:  "c",
	68:  "d",
	69:  "e",
	70:  "f",
	71:  "g",
	72:  "h",
	73:  "i",
	74:  "j",
	75:  "k",
	76:  "l",
	77:  "m",
	78:  "n",
	79:  "o",
	80:  "p",
	81:  "q",
	82:  "r",
	83:  "s",
	84:  "t",
	85:  "u",
	86:  "v",
	87:  "w",
	88:  "x",
	89:  "y",
	90:  "z",
	91:  "menu",
	92:  "menu",
	93:  "menu",
	95:  "sleep",
	96:  "numpad_0",
	97:  "numpad_1",
	98:  "numpad_2",
	99:  "numpad_3",
	100: "numpad_4",
	101: "numpad_5",
	102: "numpad_6",
	103: "numpad_7",
	104: "numpad_8",
	105: "numpad_9",
	106: "multiply",
	107: "add",
	108: "numpad period (firefox)",
	109: "subtract",
	110: "decimal point",
	111: "divide",
	112: "f1",
	113: "f2",
	114: "f3",
	115: "f4",
	116: "f5",
	117: "f6",
	118: "f7",
	119: "f8",
	120: "f9",
	121: "f10",
	122: "f11",
	123: "f12",
	124: "f13",
	125: "f14",
	126: "f15",
	127: "f16",
	128: "f17",
	129: "f18",
	130: "f19",
	131: "f20",
	132: "f21",
	133: "f22",
	134: "f23",
	135: "f24",
	136: "f25",
	137: "f26",
	138: "f27",
	139: "f28",
	140: "f29",
	141: "f30",
	142: "f31",
	143: "f32",
	144: "num lock",
	145: "scroll lock",
	151: "airplane mode",
	160: "^",
	161: "!",
	162: "؛ (arabic semicolon)",
	163: "#",
	164: "$",
	165: "ù",
	166: "page backward",
	167: "page forward",
	168: "refresh",
	169: "closing paren (AZERTY)",
	170: "*",
	171: "~ + * key",
	172: "home key",
	173: "minus (firefox), mute/unmute",
	174: "decrease volume level",
	175: "increase volume level",
	176: "next",
	177: "previous",
	178: "stop",
	179: "play/pause",
	180: "e-mail",
	181: "mute/unmute (firefox)",
	182: "decrease volume level (firefox)",
	183: "increase volume level (firefox)",
	186: "semi-colon / ñ",
	187: "equal sign",
	188: "comma",
	189: "dash",
	190: "period",
	191: "forward slash / ç",
	192: "grave accent / ñ / æ / ö",
	193: "?, / or °",
	194: "numpad period (chrome)",
	219: "open bracket",
	220: "back slash",
	221: "close bracket / å",
	222: "single quote / ø / ä",
	223: "`",
	224: "left or right ⌘ key (firefox)",
	225: "altgr",
	226: "< /git >, left back slash",
	230: "GNOME Compose Key",
	231: "ç",
	233: "XF86Forward",
	234: "XF86Back",
	235: "non-conversion",
	240: "alphanumeric",
	242: "hiragana/katakana",
	243: "half-width/full-width",
	244: "kanji",
	251: "unlock trackpad (Chrome/Edge)",
	255: "toggle touchpad}[keycode]",
}
