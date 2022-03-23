package controller

import (
	"distancing-detect-backend/internal/controller/models"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Controller) HandleAlert(client mqtt.Client, msg mqtt.Message) {
	var receivedAlert models.MqttAlert
	err := json.Unmarshal(msg.Payload(), &receivedAlert)
	if err != nil {
		c.logger.ErrorLogger.Println("Error Unmarshalling MQTT : ", err.Error())
		return
	}
	err = c.service.NewViolation(receivedAlert.Classroom, receivedAlert.TotalViolations, receivedAlert.ImageLink)
	if err != nil {
		c.logger.ErrorLogger.Println("Error Creating violations : ", err.Error())
		return
	}

}
