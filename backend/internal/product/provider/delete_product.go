package product



func DeleteProduct(c *gin.Context) {
	var product struct {
		ID int `json:"id"`
	}
	