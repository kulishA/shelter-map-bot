package message

import (
	"fmt"
	"shelter-map-bot/proto"
)

func GetHelpMessage() string {
	return "Я можу допомогти вам знайти найближчі бомбосховища, просто надішліть ваше місцезнаходження 📍"
}

func GetStartMessage() string {
	return "Будь ласка, надішліть своє місцезнаходження 📍"
}

func GetErrorMessage() string {
	return "Вибачте, але під час пошуку сталася помилка 😣 , спробуйте ще раз 🙏"
}

func GetShelterDescriptionMessage(shelter *proto.ShelterResponse) string {
	return fmt.Sprintf("📍 Адреса: %s\n"+
		"🏠 Тип укриття: %s\n"+
		"🏠 Тип будівлі: %s\n"+
		"👨🏻 Власник: %s\n"+
		"♿ Рампа для інвалідного візка: %s\n"+
		"📞 Номер телефону: %s\n",
		fmt.Sprintf("%s %s", shelter.Address, shelter.AddressNumber),
		shelter.ShelterType, shelter.BuildingType, shelter.Owner, shelter.Ramp,
		shelter.Phone,
	)
}
