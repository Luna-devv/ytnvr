[![](https://img.shields.io/discord/828676951023550495?color=5865F2&logo=discord&logoColor=white)](https://lunish.nl/support)
![](https://img.shields.io/github/repo-size/Luna-devv/nekostic?maxAge=3600)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/I3I6AFVAP)

**⚠️ In development, breaking changes ⚠️**

## About
YTNVR is a easy way to get the newest youtube video of any given youtube channel.
```bash
curl http://100.65.0.2:8080?channel_id=UClWBeVcz5LUmcCN1gHG_GCg
{"channelUrl":"https://www.youtube.com/channel/UClWBeVcz5LUmcCN1gHG_GCg","videoUrl":"https://www.youtube.com/watch?v=NS5fZ1ltovE"}
```
Since you need a api key from google to get the videos of a channel, I don't recommend to expose your api to the public, even if it implements caching.

If you need help using this, join **[our Discord Server](https://discord.com/invite/yYd6YKHQZH)**.

## Setup

To obtain an API key for the YouTube Data API, you need to follow these steps:

1. **Go to the Google Cloud Console**: Navigate to the Google Cloud Console at https://console.cloud.google.com/.
2. **Create a New Project**: If you don't already have a project, create a new one by clicking on the project dropdown menu at the top of the page and selecting "New Project." Follow the prompts to create the project.
3. **Enable the YouTube Data API**: Once you have a project, navigate to the API & Services → Library page. Search for "YouTube Data API" and click on it. Then click the "Enable" button.
4. **Create Credentials**: After enabling the API, navigate to the API & Services → Credentials page. Click on the "Create credentials" dropdown and select "API key." This will generate a new API key that you can use to access the YouTube Data API.
5. **Restrict API Key (Optional)**: For security purposes, you might want to restrict your API key. You can restrict it by IP address, HTTP referrers, or by API. This step is optional but recommended.
6. **Copy the API Key**: Once you've created the API key, copy it to your clipboard. You'll need to paste it into your code to authenticate your requests to the YouTube Data API.

Remember to keep your API key secure and don't share it publicly. If you're working in a team, consider using environment variables or other secure methods to manage your API keys. Additionally, be aware of the usage limits and billing associated with using the YouTube Data API to avoid unexpected charges.

Clone this repo with the following commands:

```bash
git clone https://github.com/Luna-devv/ytnvr
```

Create a `.env` file and add the following values:
```env
API_KEY=""
```

## Deploy

Since docker is the best thing that exists for managing deployments, it's the thing we use. If you don't want to use docker, install the go programing language from [go.dev](https://go.dev) and run `go run .` or however you want to run it.

To build the docker container run
```bash
docker build -t ytnvr .
```

To start the docker container (detached) run 
```bash
docker compose up -d
```