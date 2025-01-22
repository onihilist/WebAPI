# Utiliser l'image officielle de Golang
FROM golang:1.23-alpine AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY app/go.mod app/go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste des fichiers de l'application
COPY app/ .

# Construire l'application
RUN go build -o main .

# Étape finale
FROM alpine:latest

WORKDIR /app

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/main .

# Exposer le port de l'application
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./main"]