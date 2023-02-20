// Реализовать паттерн «адаптер» на любом примере

// Есть:
// 1. Интерфейс, с которым взаимодействует клиент - Target
// 2. Класс, который выполняет нужную логику, но не реализует клиентский интерфейс - Adaptee (тот, который нужно адаптировать)
// 3. Класс, который адаптирует существующий класс под необходимый интерфейс - Adapter (тот, который адаптирует)

package main

type Target interface {
	sendSmth()
}

type Adaptee struct {
}

// Реализация нужного метода
func (a *Adaptee) send() string {
	return "response"
}

// Тип Adapter встраивает в себя Adaptee
type Adapter struct {
	adaptee *Adaptee
}

// Adapter удовлетворяет необходимому интерфейсу
func (a *Adapter) sendSmth() {
	a.adaptee.send()
}

// конструктор Adapter
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee}
}
