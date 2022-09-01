package clinicaltrials

import "fmt"

type Continent struct {
	Name string
	Code string
}

func ccToContinent(cc string) (continentName string, continentCode string) {
	asia := Continent{
		Name: "Asia",
		Code: "AS",
	}

	europe := Continent{
		Name: "Europe",
		Code: "EU",
	}

	africa := Continent{
		Name: "Africa",
		Code: "AF",
	}

	antarctica := Continent{
		Name: "Antarctica",
		Code: "AN",
	}

	northAmerica := Continent{
		Name: "North America",
		Code: "NA",
	}

	southAmerica := Continent{
		Name: "South America",
		Code: "SA",
	}

	australia := Continent{
		Name: "Australia & Oceania",
		Code: "OC",
	}

	switch cc {
		case "AL","AD","AT","BE","BA","BG","BY","HR","CZ","DK","EE","FO","FI","AX","FR","DE","GI","GR","VA","HU","IS","IE","IT","LV","LI","LT","LU","MT","MC","MD","ME","NL","NO","PL","PT","RO","SM","RS","SK","SI","ES","SJ","SE","CH","UA","MK","GB","GG","JE","IM","UK":
			continentName = europe.Name
			continentCode = europe.Code
			break
		case "AF","BH","BD","BT","IO","BN","MM","KH","LK","CN","TW","CX","CC","PS","HK","IN","ID","IR","IQ","IL","JP","JO","KP","KR","KW","KG","LA","LB","MO","MY","MV","MN","OM","NP","PK","PH","TL","QA","SA","SG","VN","SY","TJ","TH","AE","TM","UZ","YE","XE","XD","XS":
			continentName = asia.Name
			continentCode = asia.Code
			break
		case "AG","BS","BB","BM","BZ","VG","CA","KY","CR","CU","DM","DO","SV","GL","GD","GP","GT","HT","HN","JM","MQ","MX","MS","AN","CW","AW","SX","BQ","NI","PA","PR","BL","KN","AI","LC","MF","PM","VC","TT","TC","US","VI":
			continentName = northAmerica.Name
			continentCode = northAmerica.Code
			break
		case "DZ","AO","BW","BI","CM","CV","CF","TD","KM","YT","CG","CD","BJ","GQ","ET","ER","DJ","GA","GM","GH","GN","CI","KE","LS","LR","LY","MG","MW","ML","MR","MU","MA","MZ","NA","NE","NG","GW","RE","RW","SH","ST","SN","SC","SL","SO","ZA","ZW","SS","EH","SD","SZ","TG","TN","UG","EG","TZ","BF","ZM":
			continentName = africa.Name
			continentCode = africa.Code
			break
		case "AQ","BV","GS","TF","HM":
			continentName = antarctica.Name
			continentCode = antarctica.Code
			break
		case "AR","BO","BR","CL","CO","EC","FK","GF","GY","PY","PE","SR","UY","VE":
			continentName = southAmerica.Name
			continentCode = southAmerica.Code
			break
		case "AS","AU","SB","CK","FJ","PF","KI","GU","NR","NC","VU","NZ","NU","NF","MP","FM","MH","PW","PG","PN","TK","TO","TV","WF","WS","XX":
			continentName = australia.Name
			continentCode = australia.Code
			break
		case "AZ","AM","CY","GE","KZ","RU","TR":
			continentName = fmt.Sprintf("%s,%s", europe.Name, asia.Name)
			continentCode = fmt.Sprintf("%s,%s", europe.Code, asia.Code)
			break
		case "UM":
			continentName = fmt.Sprintf("%s,%s", northAmerica.Name, australia.Name)
			continentCode = fmt.Sprintf("%s,%s", northAmerica.Code, australia.Code)
			break
	}

	return
}
