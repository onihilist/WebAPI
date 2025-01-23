# Utiliser l'image officielle de Golang
FROM golang:1.23-alpine

# Définir le répertoire de travail
WORKDIR /app

# Télécharger les dépendances
RUN go mod download

# Copier les fichiers go.mod et go.sum
COPY app/go.mod app/go.sum ./

# Copier le reste des fichiers de l'application
COPY app/ .

# Exposer le port de l'application
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["go", "run", "main.go"]