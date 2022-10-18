# HeadSpin Backend Media Service Challenge

Media server which accepts requests to upload audio WAV files. It then processes and transcodes the uploaded file to .mp3 format and saves it. A client can then make a request to get info about the file, get the transcoded .mp3 file, and delete the file.

# PreRequisites
1. Install Lame executable. (`brew install lame` for mac)
2. Golang

# Running
` go run main.go `

# Endpoints
1. (GET) /v0/audio/:audioid => Gets the given aduio file.
2. (GET) /v0/audio/:audioid/info => Gets info about the given aduio file.
3. (DELETE) v0/audio/:audioid => Deletes the given aduio file.
4. (POST) v0/audio/upload => converts wav audio file to mp3 and saves it.

# Supporing Info
1. Postman Collection.
2. Sample wav files.