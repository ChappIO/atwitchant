package twitch

//func (t *Integration) Authorize() http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		switch request.Method {
//		case "GET":
//			if t.token == "" {
//				log.Println("no token available, redirecting to twitch")
//				// we need to go fetch a token
//				loginUri := url.URL{
//					Scheme: "http",
//					Host:   request.Host,
//					Path:   "/",
//				}
//				log.Printf("expecting user to return to %s", loginUri.String())
//				redirectUri := url.URL{
//					Scheme: "https",
//					Host:   "id.twitch.tv",
//					Path:   "/oauth2/authorize",
//					RawQuery: url.Values{
//						"client_id":     []string{clientId},
//						"redirect_uri":  []string{loginUri.String()},
//						"response_type": []string{"token"},
//						"scope":         []string{"chat:read chat:edit channel:moderate whispers:read whispers:edit channel_editor"},
//					}.Encode(),
//				}
//				writer.WriteHeader(http.StatusNotFound)
//				json.NewEncoder(writer).Encode(map[string]string{
//					"redirectUri": redirectUri.String(),
//				})
//			} else {
//				log.Println("we already have a token")
//				_ = json.NewEncoder(writer).Encode(t.user)
//			}
//			break
//		case "POST":
//			log.Println("this looks like a token")
//			writer.WriteHeader(http.StatusOK)
//			// thank you for your token sir!
//			target := make(map[string]string)
//			_ = json.NewDecoder(request.Body).Decode(&target)
//			if token, ok := target["token"]; ok && token != "" {
//				log.Println("received a token from client!")
//				t.token = token
//				t.user = t.GetUser()
//				t.chat = &Chat{
//					api:  t,
//				}
//				if err := t.chat.Reconnect(); err != nil {
//					writer.WriteHeader(http.StatusInternalServerError)
//				} else {
//					_ = json.NewEncoder(writer).Encode(t.user)
//				}
//			} else {
//				log.Println("we seem to be missing the token... :(")
//			}
//			break
//		default:
//			writer.WriteHeader(http.StatusMethodNotAllowed)
//			break
//		}
//	})
//}

