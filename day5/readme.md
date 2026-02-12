# DevOps Fundamentals — Day 5 (Go Backend + Docker + GCP Deployment)

This session focused on **building a full backend application in Go**, containerizing it with **Docker**, and deploying it on **Google Cloud Platform (GCP) Compute Engine**.

## Project Overview

The session involved hands-on development and cloud deployment exercises:

- Building a **RESTful backend in Go** 🐹  
- Creating a **Docker image** for consistent environment packaging 🐳  
- Launching a **GCP Compute Engine instance** ☁️  
- Deploying the Dockerized application and exposing it via ports 🔌  

This demonstrates the **real-world workflow of coding → containerizing → cloud deployment** for DevOps engineers.

## Why This Matters in DevOps

- **Full-stack DevOps skills**: Combining coding, containerization, and cloud deployment  
- **Consistency and portability**: Docker ensures the app runs the same locally and in the cloud  
- **Infrastructure knowledge**: Deploying to a VM teaches instance management, ports, and firewall configuration  
- **Troubleshooting and debugging**: Hands-on deployment exposes real-world problems and solutions  

## Key Technical Concepts Demonstrated

- **Go Backend Development**: Writing RESTful APIs, routing, and business logic  
- **Docker & Containerization**: Creating images and running containers consistently  
- **GCP Compute Engine**: Launching and configuring cloud VMs  
- **Ports and Firewall Rules**: Exposing application safely on the internet  
- **Environment Variables**: Using `.env` files for configuration  

## Challenges and Solutions

| Challenge | Solution |
|-----------|---------|
| Docker image failed due to Go module errors | Verified `go.mod` dependencies and adjusted Dockerfile `WORKDIR` |
| App not accessible on GCP VM | Configured firewall rules to open port 8080 |
| Debugging runtime errors inside container | Used `docker logs` and iterative testing |

## Key Takeaways

- Building and containerizing applications in Go strengthens **backend and DevOps skills**  
- Docker enables **environment consistency**, reducing deployment issues  
- Deploying to GCP provides practical experience with **cloud infrastructure**  
- Understanding ports, firewalls, and VM configuration is essential for production-ready deployments  

## Deployment Steps

After creating your **Compute Engine VM**, follow these commands:

```bash
# Update package index
sudo apt update -y

# Install Docker
sudo apt-get install docker.io -y

# Add current user to docker group (requires logout/login to take effect)
sudo usermod -aG docker $USER

# Run the Docker container with environment variables and port mapping
docker run --name blog-api-container --env-file .env -p 8080:8080 param9999/blog-backend:latest
