package service

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func Fetch() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	page.SetExtraHTTPHeaders(map[string]string{
		":authority":                "www.singaporeair.com",
		":method":                   "GET",
		":path":                     "/en_UK/sg/home",
		":scheme":                   "https",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-encoding":           "gzip, deflate, br",
		"accept-language":           "en-US,en;q=0.9",
		"cache-control":             "max-age=0",
		"cookie":                    "saadevice=desktop; AKAMAI_SAA_DEVICE_COOKIE=desktop; AKAMAI_SAA_COUNTRY_COOKIE=SG; AKAMAI_SAA_LOCALE_COOKIE=en_UK; FARE_DEALS_LISTING_COOKIE=false; dtCookie=v_4_srv_2_sn_3664C11F584212C319667BDEBB8BA9D6_perc_100000_ol_0_mul_1_app-3A7bacf3aecd29efdf_0; hvsc=true; AKAMAI_SAA_DEVICE_COOKIE=desktop; AKAMAI_HOME_PAGE_ACCESSED=true; VSUUID=ad1fa1d553c170f2a8ce2e376ded9ee6b5610761c106ee177f3febc91d9a92e2.saa-home-8-7r8vk.1669268649228; kfMagic=K_3bd11c3b73cc5e310ebec33679afe9b1; AKA_A2=A; _abck=55834C8C6B3FE2CB1C436F111578F992~0~YAAQkvpWuFxR5YeGAQAAPVe6ogmolEZUJJwc6RVftQ/UqnGkdyfmi0hGj5gKTfaU2KZChklI4J9RrWFClqAEqDLOCC65jckaN3/cQBLKxHHkvDd2YPIKRAYO/aIbdKWYc/2UhIv+mWQOHC/d/pv+nyBoWM8tHzriTqzm2NMpYNqKOnIZhhrv5fClkBySAK5NblU08At7fxvkG1SGMW9NnLRrFzp5nqgCMgBXGePPf0OGsXpFZfcnufQYZ2XuKEXPSem5feQTMUFnodQ8Hjp3cLvnKHmbJaud/39IxB9oe/oVH3VL2haN8hlW6xbVcFY5cjXyDvn83ldpjFTIDLckc2L8zPGATcxTyeWsZ3qLHo0lwhdmNwCIV/3/YitB2WbHuihQ5Zvy88JkdctHPtvuehncyp1XlmB32x2AlZNR~-1~-1~-1; ak_bmsc=7F0992FB177C8BB6C77CB9BD3C65EFBA~000000000000000000000000000000~YAAQkvpWuF1R5YeGAQAAPVe6ohLz8YUSSNrY8pgqFPNxkR1q1mQgB2irHuq8ttHnrrZLjHICU0WFkfD0HGJPiCN484ENCUCdtw7FuxcjdILNTkUxySt7OOqIg46iks0lgHCCaJLh7HgMFrTeg51z1prI/7tk7QcluyYvA4fiuY3gMk06cKNNdM9We1HdAuDZzMp06uO6kfgQlVroNxx9EgvRXu3ESPKF7vJY4TPc5J7GVqCS4lEpOg0BFmuaEFFFVInBtSv+sHLzCV/3rKb+ZrdPSDoZMqT95gS0VYiZlf2ZINOO1N9XEp+fWJ0k0I4CLg/vZB6lIlUiZ+scwihswuJ3ZoDquazxqHwIMzWtyJpyCT+VC1UIVWswLuxm51DU1vOl3roPvja4BIW8RYTerSYy; bm_sz=0CCE4E214438246B47D5FA87284FFB51~YAAQkvpWuF5R5YeGAQAAPVe6ohIcYkxZAC0yIlttQRsjPCO6tWM/WXrZkfVWlwfvKXFD78s2CphyXkaG18NbBrglSNLH+GZcAj8995peX+CbME9S7Ceq4uemOOyn4vDIRgB4NoukpmNlHEM7FlO4zSsKPeoouU2EbnEGdm/3UMEg960daQqIgvLCGXQGlOl+o65DzAnQH7uavx1Cnb5UGkcTO9ztd6i2zcffzizw7N2S6PbkNMAIY1KYCQmGT9iOg5Tjg2M0me3/saRmr+Fk5CqkHQNoxa/8roaBBzpWQjwJ+Edgy8IS5UQ=~4534841~4273209; AKAMAI_SAA_AIRPORT_COOKIE=SIN; HSESSIONID=DLkz0t1xIlkdnmCcznKBfxo13DLicZ8VF08pRWH3.saa-home-13-x4q29; LOGIN_COOKIE=false; LOGIN_POPUP_COOKIE=false; RU_LOGIN_COOKIE=false; SQCLOGIN_COOKIE=false; 6b29450cab647be0f08ef134c7afc9a1=9b360b4183581d5ddd94ebb887cd15c3; bm_mi=2D1A4CBE54607CF0E0DE9251B8DE948C~YAAQkvpWuGRR5YeGAQAAxle6ohLB+FIx3Wgw1P9VniA9IJpqnJ01oO/de8ldVMMaXAmUaISKOOciCUe84fpnuRSfL29qT980A3QQk+mrXypsynLvha6f36UapXnRz7siHeEdCWzSsLrzUj7IefAkokwuC1xLWN7zBx+o9Ii1apa4N/ETRPg8S6iCjyM8SIMOqbe6n49+pXF0YFamuIQ+prV83pt2JZ2gLSmd0q5w0bzMOTs4O+6fl1ZYiVf/fbrLKoLw91z3vYbTrVJU4W3KH7g9dw0/F3WJvRRKndhRU6AzxPjANvmDz2ojuPUmYvMGgWX6stHibKC944NVksXecQ==~1; AWSALB=LPKDj57QYSnFSEfRSbkPkOgxRMD0uQCjM/UoKiOyteF4ofs5sYy1p16dW5XRTW9C8vQ5zjEsaMSbrl554xnr+kaGjBEbmZ7nsHTBYrWLUU0UeQiBf0+4j9K1T5BA; AWSALBCORS=LPKDj57QYSnFSEfRSbkPkOgxRMD0uQCjM/UoKiOyteF4ofs5sYy1p16dW5XRTW9C8vQ5zjEsaMSbrl554xnr+kaGjBEbmZ7nsHTBYrWLUU0UeQiBf0+4j9K1T5BA; JSESSIONID=xMqiul3KmoOsxPQgk7MVT523wdwcsu7s5w5b-DUOEoDNJULFBKj0!-749651164; cookieStateSet=hybridState; bm_sv=87080619C526662FCDBCA538E839E875~YAAQkvpWuLxR5YeGAQAA1me6ohLIWsTjVtrdS68HvT3V3te61KdZupKHF4CJQs2Bi2XIWzfHF9Ew1DdeR33ARyeE7P48VXsHIttcSX97BnXG6dGmmyBbDKQtOez5r9qtnfIBxf4zfY75BL7mczlnTmnv+FJAzerxm3scHQ10ps803D8g3iI2vOQVq0zXpmAZmnm7NZYZSXBvQjwPvkPh6PCvk64uav6s/D5YHKXIQ49YncvjJ1sDs6ggBazV37FZNoI83fkn~1",
		"referer":                   "https://www.google.com/",
		"sec-ch-ua":                 "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"macOS\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/111.0",
	})
	if _, err = page.Goto("https://www.singaporeair.com/en_UK/sg/home#/book/bookflight"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	loginbutton, err := page.Locator("text=Log-in")
	assertErrorToNilf("could not click on login button: %v", err)
	loginbutton.Click()

	userinput, err := page.Locator("xpath=//html/body/div[1]/div[4]/sia-header/div[1]/light-box[1]/div/div/div/sia-login/context-consumer/div/div[4]/div/input-box[1]/div/div/label/input")
	assertErrorToNilf("could not find user input", err)
	userinput.Type("nick.chow.zj@gmail.com")

	passwordinput, err := page.Locator("xpath=//html/body/div[1]/div[4]/sia-header/div[1]/light-box[1]/div/div/div/sia-login/context-consumer/div/div[4]/div/input-box[2]/div/div/label/input")
	assertErrorToNilf("could not find password input", err)
	passwordinput.Type("r0landgileaD")

	loginbutton, err = page.Locator("text=Log in")
	assertErrorToNilf("could not click on login button 2: %v", err)
	loginbutton.Click()

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}

	//assertErrorToNilf("could not click on login button 2: %v", page.Click(".dwc--SiaLogin__ButtonLogin"))
}
