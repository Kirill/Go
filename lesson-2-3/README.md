.center.icon[![otus main](img/main.png)]

---

class: top white
background-image: url(img/sound.svg)
background-size: 130%
.top.icon[![otus main](img/logo.png)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **-** если есть проблемы cо звуком или с видео
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Форматирование данных

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

### 1. Зачем нужна сериализация/десериализация.
### 2. Текстовая сериализация (json, xml).
### 3. Бинарная сериализация (msgpack, protobuf).

---

# Настройка на занятие

.left-text[
Пожалуйста, пройдите небольшой тест.
<br><br>
Возможно, вы уже многое знаете про возможности сериализации в Go.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher9.png)
]

---

# Зачем?

### Зачем кодировать бинарные данные в текстовый вид?

---

# base64

```
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
```

```
MS4g0JfQsNGH0LXQvCDQvdGD0LbQvdCwINGB0LXRgNC40LDQu9C40LfQsNGG0LjRjy/QtNC
10YHQtdGA0LjQsNC70LjQt9Cw0YbQuNGPLgoyLiDQotC10LrRgdGC0L7QstCw0Y8g0YHQtd
GA0LjQsNC70LjQt9Cw0YbQuNGPIChqc29uLCB4bWwpLgozLiDQkdC40L3QsNGA0L3QsNGPI
NGB0LXRgNC40LDQu9C40LfQsNGG0LjRjyAobXNncGFjaywgcHJvdG9idWYpLg==
```

Избыточность = 1/3

---

# base64
```
func main() {
    data := "Hello world"

    sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
}
```

https://play.golang.org/p/eoIUTskNcMf

---

# base64: поточная работа

Кодирование
```
func main() {
	input := []byte("otus")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close() // <=== Зачем?
}
```
<br>
Декодирование
```
func main() {
	input := "b3R1cw=="
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
	io.Copy(os.Stdout, r)
}
```

---

# base64: недостатки

Какие недостатки?

---

# base64: недостатки

```
Std: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
URL: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
```

---

# Текстовая сериализация
## JSON
## XML
## YAML

Какие цели у сериализации?

---

# JSON

```
{"widget": {
    "debug": "on",
    "window": {
        "title": "Sample Konfabulator Widget",
        "name": "main_window",
        "width": 500,
        "height": 500
    },
    "image": {
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    },
    "text": {
        "data": "Click Here",
        "size": 36,
        "style": "bold",
        "name": "text1",
        "hOffset": 250,
        "vOffset": 100,
        "alignment": "center",
        "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
}}
```

---

# YAML

```
---
widget:
  debug: 'on'
  window:
    title: Sample Konfabulator Widget
    name: main_window
    width: 500
    height: 500
  image:
    src: Images/Sun.png
    name: sun1
    hOffset: 250
    vOffset: 250
    alignment: center
  text:
    data: Click Here
    size: 36
    style: bold
    name: text1
    hOffset: 250
    vOffset: 100
    alignment: center
    onMouseUp: sun1.opacity = (sun1.opacity / 100) * 90;
```

---

# XML

```
<widget>
    <debug>on</debug>
    <window title="Sample Konfabulator Widget">
        <name>main_window</name>
        <width>500</width>
        <height>500</height>
    </window>
    <image src="Images/Sun.png" name="sun1">
        <hOffset>250</hOffset>
        <vOffset>250</vOffset>
        <alignment>center</alignment>
    </image>
    <text data="Click Here" size="36" style="bold">
        <name>text1</name>
        <hOffset>250</hOffset>
        <vOffset>100</vOffset>
        <alignment>center</alignment>
        <onMouseUp>
            sun1.opacity = (sun1.opacity / 100) * 90;
        </onMouseUp>
    </text>
</widget>
```

---

# JSON: структуры

```
func main() {
	p1 := &Person{
		Name: "Vasya",
		Age: 36,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("p1 json %s\n", j)

	var p2 Person
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %+v\n", p2)
}
```

https://play.golang.org/p/c6EmkEXVl1_a

---

# JSON: interface{}

```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := []byte(`{"Name":"Vasya",
		"Job":{"Department":"Operations","Title":"Boss"}}`)

	var p2 interface{}
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

	person := p2.(map[string]interface{})
	fmt.Printf("name=%s\n", person["Name"])
}
```

https://play.golang.org/p/_4tBO2EgDVE

---

# JSON: тэги

```
type Person struct {
	Name    string `json:"fullname,omitempty"`
	Surname string `json:"familyname,omitempty"`
	Age     int    `json:"-"`
	Job     struct {
		Department string
		Title      string
	}
}
```

---

# JSON: easyjson

https://github.com/mailru/easyjson

---

# XML

```
type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}
```

https://play.golang.org/p/QbfwL44vjJU (сериализация)


https://play.golang.org/p/FekJkpuj9KT (десериализация)

---

# XML: SAX Parser

```
for {
		token, _ := decoder.Token()

		switch se := token.(type) {
		case xml.StartElement:
			fmt.Printf("Start element: %v Attr %s\n",
						se.Name.Local, se.Attr)
			inFullName = se.Name.Local == "FullName"
		case xml.EndElement:
			fmt.Printf("End element: %v\n", se.Name.Local)
			inFullName = false
		case xml.CharData:
			fmt.Printf("Data element: %v\n", string(se))
			if inFullName {
				names = append(names, string(se))
			}
		default:
			fmt.Printf("Unhandled element: %T", se)
		}

	}
```

https://play.golang.org/p/5qU9TYTo-G9

---

# Бинарные сериализаторы

## Что это такое?
## Какие плюсы?
## Какие минусы?

---

# Бинарные сериализаторы

## gob
## msgpack
## protobuf

---

# msgpack

```
type Person struct {
	Name        string
	Surname     string
	Age         uint8
	ChildrenAge map[string]uint8
}
func main() {
	p := Person{Name:  "Rob",
		Surname: "Pike", Age: 27,
	}
	p.ChildrenAge = make(map[string]uint8)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, _ := msgpack.Marshal(&p)

	fmt.Printf("Length of marshaled: %v\n", len(marshaled))
	fmt.Printf("Binary: %v\n", string(marshaled))

	p2 := &Person{}
	msgpack.Unmarshal(marshaled, p2)
	fmt.Printf("Unmarshled: %+v\n", p2)
}
```

https://play.golang.org/p/VrLsTxv68V_O

---

# protobuf

```
syntax = "proto3";

package main;

message Person {
    string name = 1;
    string surname = 2;
    uint32 age = 3;

    map<string, uint32> children_age = 4;
}
```

Кодогенерация: `protoc --go_out=. *.proto`.

---

# protobuf

```
func main() {
	p := &Person{
		Age:         27,
		Name:        "Rob",
		Surname:     "Pike",
		ChildrenAge: make(map[string]uint32),
	}
	p.ChildrenAge["Maria"] = 2
	p.ChildrenAge["Alex"] = 5

	marshaled, _ := proto.Marshal(p)

	fmt.Printf("Length of marshaled: %v\n", len(marshaled))
	fmt.Printf("Binary: %v\n", string(marshaled))

	p2 := &Person{}
	proto.Unmarshal(marshaled, p2)

	fmt.Printf("Unmarshaled %v\n", p2)
}
```

---

# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher9.png)
]

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате.
]

.right-image[
![](img/gopher.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
