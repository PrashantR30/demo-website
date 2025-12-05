package main

import (
    "crypto/md5"            // ❌ insecure hashing
    "crypto/sha1"           // ❌ insecure hashing
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/exec"               // ❌ command injection
    "database/sql"          // ❌ sql injection
    _ "github.com/go-sql-driver/mysql"
)

var hardcodedAPIKey = "AIzaSyA123456789ABCDEFG" // ❌ #1 Hardcoded secret

// ===============================================================
// #1 Hardcoded Credential     (high)
// ===============================================================

var hardcodedDBPass = "rootpassword123" // ❌ Hardcoded DB password

// ===============================================================
// #2 Command Injection        (high)
// ===============================================================

func vulnCommandInjection(w http.ResponseWriter, r *http.Request) {
    userInput := r.URL.Query().Get("cmd")
    out, _ := exec.Command("sh", "-c", userInput).CombinedOutput() // ❌ vulnerable
    w.Write(out)
}

// ===============================================================
// #3 SQL Injection            (high)
// ===============================================================

func vulnSQLInjection(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("user")

    query := "SELECT * FROM users WHERE username = '" + username + "'" // ❌ vulnerable

    db, _ := sql.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    db.Query(query)
}

// ===============================================================
// #4 Path Traversal           (high)
// ===============================================================

func vulnPathTraversal(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    data, _ := ioutil.ReadFile("/var/data/" + file) // ❌ vulnerable
    w.Write(data)
}

// ===============================================================
// #5 Unsafe JSON Unmarshal    (medium)
// ===============================================================

type Profile struct {
    Name string `json:"name"`
}

func vulnUnsafeUnmarshal(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    var p Profile
    json.Unmarshal(body, &p) // ❌ no error check
    fmt.Fprintf(w, "Hello %s", p.Name)
}

// ===============================================================
// #6 Insecure Randomness       (medium)
// ===============================================================

func vulnInsecureRandomness() {
    b := make([]byte, 16)
    // ❌ insecure RNG used for tokens
    f, _ := os.Open("/dev/urandom") // Semgrep flags this in some rulesets
    f.Read(b)
    fmt.Println(b)
}

// ===============================================================
// #7 Insecure Hashing (md5)    (high)
// ===============================================================

func vulnMD5() {
    data := []byte("password123")
    _ = md5.Sum(data) // ❌ insecure hash
}

// ===============================================================
// #8 Insecure Hashing (sha1)   (high)
// ===============================================================

func vulnSHA1() {
    data := []byte("password123")
    _ = sha1.Sum(data) // ❌ insecure hash
}

// ===============================================================
// #9 Unsafe Temp File          (medium)
// ===============================================================

func vulnUnsafeTempFile() {
    f, _ := os.Create("/tmp/hardcoded_temp_file") // ❌ predictable temp file
    defer f.Close()
    f.WriteString("test")
}

// ===============================================================
// #10 JWT NONE Algorithm        (critical)
// ===============================================================

func vulnJWTNoneAlgorithm() {
    header := `{"alg":"none","typ":"JWT"}` // ❌ semgrep detects insecure JWT usage
    payload := `{"user":"admin"}`
    token := fmt.Sprintf("%s.%s.", header, payload) // ❌ unsigned JWT
    fmt.Println(token)
}

// dummy main
func main() {}

