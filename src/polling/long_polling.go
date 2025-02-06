package polling

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var longPollingCh = make(chan struct{})

// Función que simula el envío de actualizaciones al canal
func simulateUpdates() {
	for {
		time.Sleep(10 * time.Second) // Simulamos una actualización cada 10 segundos
		longPollingCh <- struct{}{}   // Envía la actualización al canal
	}
}

// Endpoint de long polling
func LongPolling(c *gin.Context) {
	select {
	case <-longPollingCh:
		c.JSON(http.StatusOK, gin.H{"update": "Product updated"})
	case <-time.After(30 * time.Second): // Tiempo de espera
		c.JSON(http.StatusOK, gin.H{"message": "No updates available"})
	}
}
