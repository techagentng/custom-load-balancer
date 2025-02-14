# **Gin Load Balancer** 🚀

A simple **custom load balancer** built with **Go (Gin)** to distribute traffic across multiple backend services using **round-robin**.

## **Features**  
✅ Reverse proxy for forwarding requests  
✅ Round-robin load balancing  
✅ Supports multiple backend instances  

## **Getting Started**  

### **1. Clone the repo**  
```bash
git clone https://github.com/your-username/gin-load-balancer.git
cd gin-load-balancer
```

### **2. Install dependencies**  
```bash
go mod tidy
```

### **3. Run the load balancer**  
```bash
go run main.go
```

### **4. Start backend servers**  
Run multiple backend instances on different ports (e.g., 8081, 8082, 8083).  

### **5. Test the load balancer**  
Send requests to:  
```bash
curl http://localhost:8080/some-endpoint
```
Traffic will be distributed across backends! 🎯  

## **Contributing**  
Feel free to fork, improve, and open a PR! 🚀  

## **License**  
MIT 📜  


