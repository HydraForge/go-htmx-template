templ:
	@templ generate -watch -proxy=http://localhost:8000

tailwind: 
	@npx tailwindcss build -o src/assets/css/main.css --watch