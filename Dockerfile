# Définir l'image de base
FROM golang:latest

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers du projet
COPY . .

# Installer git, nécessaire pour fetch les dépendances
RUN apt-get update && apt-get install -y git

# Télécharger les dépendances
RUN go mod download

# Construire l'application
RUN go build -o main .

# Exécuter l'application
CMD ["/app/main"]