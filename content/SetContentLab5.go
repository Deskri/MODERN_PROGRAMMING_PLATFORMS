package content

import (
	"MODERN_PROGRAMMING_PLATFORMS/work"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetContentLab5(window fyne.Window) {
	//textInfoTask1 := "Задание 1. "
	//textInfoTask2 := "Задание 2. "
	//textInfoTask3 := "Задание 3. "
	//textInfoTask4 := "Задание 4. "
	textInfoTask5 := "Задание 5. Выведите все четные элементы списка. При этом используйте цикл for, перебирающий элементы списка, а не их индексы!"
	//textInfoTask6 := "Задание 6. "
	//textInfoTask7 := "Задание 7. "
	//textInfoTask8 := "Задание 8. "

	//ФОРМА ЗАДАНИЕ 5
	// поле ввода
	entryTask5 := widget.NewEntry()
	// лейбл вывода ответа
	labelAnswerTask5 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask5 := widget.NewForm(
		widget.NewFormItem("Введите массив 5-ти чисел разделенные пробелом:", entryTask5),
		widget.NewFormItem("Ответ:", labelAnswerTask5),
	)
	// присваивание кнопки названия "Выполнить"
	formTask5.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask5.OnSubmit = func() {
		labelAnswerTask5.SetText(work.Lab5Task5(entryTask5.Text))
	}

	//accordionItemTask1 := widget.NewAccordionItem(
	//	textInfoTask1,
	//	formTask1,
	//)
	//accordionItemTask2 := widget.NewAccordionItem(
	//	textInfoTask2,
	//	formTask2,
	//)
	//accordionItemTask3 := widget.NewAccordionItem(
	//	textInfoTask3,
	//	formTask3,
	//)
	//accordionItemTask4 := widget.NewAccordionItem(
	//	textInfoTask4,
	//	formTask4,
	//)
	accordionItemTask5 := widget.NewAccordionItem(
		textInfoTask5,
		formTask5,
	)
	//accordionItemTask6 := widget.NewAccordionItem(
	//	textInfoTask6,
	//	formTask6,
	//)
	//accordionItemTask7 := widget.NewAccordionItem(
	//	textInfoTask7,
	//	formTask7,
	//)
	//accordionItemTask8 := widget.NewAccordionItem(
	//	textInfoTask8,
	//	formTask8,
	//)
	accordion := widget.NewAccordion(
	//	accordionItemTask1,
	//	accordionItemTask2,
	//	accordionItemTask3,
	//	accordionItemTask4,
		accordionItemTask5,
	//	accordionItemTask6,
	//	accordionItemTask7,
	//	accordionItemTask8,
	)

	window.SetContent(container.NewVScroll(container.NewVBox(
		accordion)))
}