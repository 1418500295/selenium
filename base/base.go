package base

import "github.com/tebeka/selenium"

type Base struct {
	Driver selenium.WebDriver
}

func (base *Base) driver() selenium.WebDriver {
	return base.Driver
}

func (base *Base) FindEleByName(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByName, element)
	return ele
}
func (base *Base) FindEleByClassName(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByClassName, element)
	return ele
}
func (base *Base) FindEleByID(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByID, element)
	return ele
}
func (base *Base) FindEleByCssSelector(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByCSSSelector, element)
	return ele
}
func (base *Base) FindEleByXpath(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByXPATH, element)
	return ele
}
func (base *Base) FindEleByLinkText(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByLinkText, element)
	return ele
}

func (base *Base) FindEleByTagName(element string) selenium.WebElement {
	ele, _ := base.Driver.FindElement(selenium.ByTagName, element)
	return ele
}

func (base *Base) FindElsByName(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByName, element)
	return ele
}
func (base *Base) FindElsByClassName(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByClassName, element)
	return ele
}
func (base *Base) FindElsByID(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByID, element)
	return ele
}
func (base *Base) FindElsByCssSelector(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByCSSSelector, element)
	return ele
}
func (base *Base) FindElsByXpath(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByXPATH, element)
	return ele
}
func (base *Base) FindElsByLinkText(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByLinkText, element)
	return ele
}

func (base *Base) FindElsByTagName(element string) []selenium.WebElement {
	ele, _ := base.Driver.FindElements(selenium.ByTagName, element)
	return ele
}
