package info

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (text string, err error) {
	text = "[VERSION 0.1.1]\n" +
		"> sup ptr\t123.45.67.89\n" +
		"> sup rps\trequests since until (ex.: sup rps 12345 '2023-03-30 11:00:00' '2023-03-30 11:07:00')\n" +
		"> sup 443|80\tdomain target_ip (ex.: sup 443 example.com 93.184.216.34)\n" +
		"> sup hosts\tdomain target_ip (as sudo)\n" +
		"> sup har\tpath_to_har_file target_ip echo_ip (for any Chromium)\n" +
		"> sup ea\t(.bashrc must include: LDAPUSER, ECHO_FILE, ECHO_LIST(', '))"
	return
}
