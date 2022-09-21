package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	defer func() {
		if err = browser.Close(); err != nil {
			log.Printf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Printf("could not stop Playwright: %v", err)
		}
	}()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("http://adactinhotelapp.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	err = fillLoginPage(page)
	if err != nil {
		log.Fatalf("fialed filling the login page, error : %v", err)
	}

	err = submitLoginPage(page)
	if err != nil {
		log.Fatalf("fialed submitting the login page, error : %v", err)
	}

	page.WaitForNavigation()

	err = fillBookingPage(page)
	if err != nil {
		log.Fatalf("fialed filling the booking page, error : %v", err)
	}

	err = submitBookingPage(page)
	if err != nil {
		log.Fatalf("failing submitting the booking page, error %v", err)
	}

	page.WaitForNavigation()

	err = fillConfirmationPage(page)
	if err != nil {
		log.Fatalf("failed filling the confirmation page, error %v", err)
	}

	err = submitConfirmationPage(page)
	if err != nil {
		log.Fatalf("failed submitting the confirmation page, error %v", err)
	}
	page.WaitForNavigation()

	err = fillDetailsPage(page)
	if err != nil {
		log.Fatalf("failed filling the details page, error %v", err)
	}

	err = confirmBooking(page)
	if err != nil {
		log.Fatalf("failed submitting the booking confirmation, error %v", err)
	}
	page.WaitForNavigation()

	fullPage := true
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("two2.png"), FullPage: &fullPage}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	log.Println("4th page Screenshot created")
}

// TODO: add submit details function

func fillDetailsPage(page playwright.Page) error {
	elFirstname, err := page.QuerySelector("#first_name")
	if err != nil {
		return fmt.Errorf("could not get firstname: %v", err)
	}
	err = elFirstname.Fill("Macruz")
	if err != nil {
		return fmt.Errorf("could not get name: %v", err)
	}

	elLastname, err := page.QuerySelector("#last_name")
	if err != nil {
		return fmt.Errorf("could not get lastname: %v", err)
	}
	err = elLastname.Fill("John")
	if err != nil {
		return fmt.Errorf("could not get last_name: %v", err)
	}

	elAddress, err := page.QuerySelector("#address")
	if err != nil {
		return fmt.Errorf("could not get address: %v", err)
	}
	err = elAddress.Fill("Chennai")
	if err != nil {
		return fmt.Errorf("could not get address: %v", err)
	}

	elCcNumber, err := page.QuerySelector("#cc_num")
	if err != nil {
		return fmt.Errorf("could not get cc_num loc: %v", err)
	}
	err = elCcNumber.Fill("9876543298765432")
	if err != nil {
		return fmt.Errorf("could not get cc_num: %v", err)
	}

	elCardType, err := page.QuerySelector("#cc_type")
	if err != nil {
		return fmt.Errorf("could not find cardtype: %v", err)
	}

	cardType := playwright.SelectOptionValues{
		Values: &[]string{"VISA"},
	}
	_, err = elCardType.SelectOption(cardType)
	if err != nil {
		return fmt.Errorf("could not select card type: %v", err)
	}

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("two3.png")}); err != nil {
		return fmt.Errorf("could not create screenshot: %v", err)
	}
	log.Println("3 page Screenshot created")

	elMonth, err := page.QuerySelector("#cc_exp_month")
	if err != nil {
		return fmt.Errorf("could not find month: %v", err)
	}
	month := playwright.SelectOptionValues{
		Values: &[]string{"3"},
	}
	_, err = elMonth.SelectOption(month)
	if err != nil {
		return fmt.Errorf("could not select month: %v", err)
	}

	elYear, err := page.QuerySelector("#cc_exp_year")
	if err != nil {
		return fmt.Errorf("could not find year: %v", err)
	}
	year := playwright.SelectOptionValues{
		Values: &[]string{"2022"},
	}
	_, err = elYear.SelectOption(year)
	if err != nil {
		return fmt.Errorf("could not select year: %v", err)
	}

	elCcv, err := page.QuerySelector("#cc_cvv")
	if err != nil {
		return fmt.Errorf("could not get cc_cvv loc: %v", err)
	}
	err = elCcv.Fill("432")
	if err != nil {
		return fmt.Errorf("could not get cc_cvv: %v", err)
	}
	return nil
}

func fillLoginPage(page playwright.Page) error {
	elUsername, err := page.QuerySelector("#username")
	if err != nil {
		return fmt.Errorf("could not get username: %v", err)
	}
	err = elUsername.Fill("Macruz123")
	if err != nil {
		return fmt.Errorf("could not get username: %v", err)
	}

	elPassword, err := page.QuerySelector("#password")
	if err != nil {
		return fmt.Errorf("could not get password: %v", err)
	}
	err = elPassword.Fill("9876543210")
	if err != nil {
		return fmt.Errorf("could not get username: %v", err)
	}

	return nil
}

func submitLoginPage(page playwright.Page) error {
	elLoginBtn, err := page.QuerySelector("#login")
	if err != nil {
		return fmt.Errorf("could not get login: %v", err)
	}
	err = elLoginBtn.Click()
	if err != nil {
		return fmt.Errorf("could not click login: %v", err)
	}

	return nil
}

func fillBookingPage(page playwright.Page) error {
	_, err := page.QuerySelector("#search_form")
	if err != nil {
		return fmt.Errorf("could not find search form: %v", err)
	}
	log.Println("moved to search form after login")

	elLocation, err := page.QuerySelector("#location")
	if err != nil {
		return fmt.Errorf("could not find location dropdown: %v", err)
	}
	location := playwright.SelectOptionValues{
		Values: &[]string{"Sydney"},
	}
	_, err = elLocation.SelectOption(location)
	if err != nil {
		return fmt.Errorf("could not select location dropdown: %v", err)
	}
	log.Println("location selected")

	elHotels, err := page.QuerySelector("#hotels")
	if err != nil {
		return fmt.Errorf("could not find hotels dropdown: %v", err)
	}
	hotel := playwright.SelectOptionValues{
		Values: &[]string{"Hotel Creek"},
	}
	_, err = elHotels.SelectOption(hotel)
	if err != nil {
		return fmt.Errorf("could not select hotels dropdown: %v", err)
	}

	elRoomType, err := page.QuerySelector("#room_type")
	if err != nil {
		return fmt.Errorf("could not find room type dropdown: %v", err)
	}

	roomType := playwright.SelectOptionValues{
		Values: &[]string{"Deluxe"},
	}
	_, err = elRoomType.SelectOption(roomType)
	if err != nil {
		return fmt.Errorf("could not select room type dropdown: %v", err)
	}

	elRoomNos, err := page.QuerySelector("#room_nos")
	if err != nil {
		return fmt.Errorf("could not find room nos dropdown: %v", err)
	}

	roomNos := playwright.SelectOptionValues{
		Values: &[]string{"1"},
	}
	_, err = elRoomNos.SelectOption(roomNos)
	if err != nil {
		return fmt.Errorf("could not select room nos dropdown: %v", err)
	}

	elAdultRoom, err := page.QuerySelector("#adult_room")
	if err != nil {
		return fmt.Errorf("could not find room nos dropdown: %v", err)
	}

	adultRoom := playwright.SelectOptionValues{
		Values: &[]string{"2"},
	}
	_, err = elAdultRoom.SelectOption(adultRoom)
	if err != nil {
		return fmt.Errorf("could not select room nos dropdown: %v", err)
	}
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("two1.png")}); err != nil {
		return fmt.Errorf("could not create screenshot: %v", err)
	}
	log.Println("Screenshot created")
	return nil
}

func submitBookingPage(page playwright.Page) error {
	elSubmitBtn, err := page.QuerySelector("#Submit")
	if err != nil {
		return fmt.Errorf("could not get submit: %v", err)
	}
	err = elSubmitBtn.Click()
	if err != nil {
		return fmt.Errorf("could not click submit: %v", err)
	}
	return nil
}

func fillConfirmationPage(page playwright.Page) error {
	elRadioBtn, err := page.QuerySelector("#radiobutton_0")
	if err != nil {
		return fmt.Errorf("could not get radio btn: %v", err)
	}
	err = elRadioBtn.Click()
	if err != nil {
		return fmt.Errorf("could not click radio btn: %v", err)
	}

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("two.png")}); err != nil {
		return fmt.Errorf("could not create screenshot: %v", err)
	}
	log.Println("2 page Screenshot created")

	return nil
}
func submitConfirmationPage(page playwright.Page) error {
	elContBtn, err := page.QuerySelector("#continue")
	if err != nil {
		return fmt.Errorf("could not get continue: %v", err)
	}
	err = elContBtn.Click()
	if err != nil {
		return fmt.Errorf("could not click continue btn: %v", err)
	}

	return nil
}

func confirmBooking(page playwright.Page) error {
	elBookingbtn, err := page.QuerySelector("#book_now")
	if err != nil {
		return fmt.Errorf("could not get login: %v", err)
	}
	err = elBookingbtn.Click()
	if err != nil {
		return fmt.Errorf("could not click login: %v", err)
	}

	return nil
}
