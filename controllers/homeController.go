package controllers

import (
	"blog/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	navData := services.GetNavs()
	fmt.Println(navData.SiteName.PartOne)
	return c.Render("home", fiber.Map{
		"NavData": navData,
	})
}
