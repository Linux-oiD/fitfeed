package main

import (
 "fitfeed/auth/utils"
 "fitfeed/auth/models"
)

func init() {
 utils.LoadEnvVariables()
 utils.ConnectToDB()
}
