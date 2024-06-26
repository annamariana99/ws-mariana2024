package url

import (
	"github.com/annamariana99/ws-mariana2024/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	page.Get("/presensi", controller.GetPresensi)
	//di buat untuk memanggil url buat get by idnya 
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
