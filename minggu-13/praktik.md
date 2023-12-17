# Praktikum Teknologi Cloud Minggu ke-13
## Hello Minikube
Praktikum ini membahas mengenai cara menjalankan aplikasi di Kubernetes menggunakan minikube.  
  
### Create a minikube cluster
Jalankan minikube dengan perintah:  
```
$ minikube start
```
![menjalankan minikube dengan perintah minikube start pada terminal](./01-minikube-start.png)  
  
Buka dasbor kubernetes dengan perintah berikut:  
```
$ minikube dashboard --url
```
![menjalankan dasbor minikube dengan perintah minikube dashboard --url](./02-minikube-dashboard.png)  
![tampilan dasbor minikube](./03-minikube-dashboard.png)  
  
Buat sebuah *Deployment* untuk memanajemen *Pod* dengan menggunakan perintah berikut:  
```
$ kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080
```
![membuat sebuah deployment dengan perintah kubbectl create deployment](./04-kubectl-create-deployment.png)  
  
Jalankan perintah `kubectl get deployments` untuk melihat *Deployment* yang sudah dibuat.  
![melihat deployment menggunakan perintah kubectl get deployments](./05-kubectl-get-deployments.png)  
  
Jalankan perintah `kubectl get pods` untuk melihat daftar *Pod* yang berjalan. 
![melihat daftar pod dengan menggunakan perintah kubectl get pods](./06-kubectl-get-pods.png)   
  
Jalankan perintah `kubectl get events` untuk melihat *cluster events*.  
![melihat events dengan perintah kubectl get events](./07-kubectl-get-events.png)  
  
Jalankan perintah `kubectl config view` untuk melihat konfigurasi.  
![melihat konfigurasi dengan perintah kubectl config view](./08-kubectl-config-view.png)  
  
Jalnakan perintah `kubectl logs hello-node-ccf4b9788-p7bc6` untuk melihat log dari aplikasi yang sedang berjalan.  
![melihat log aplikasi dengan perintah kubectl logs](./09-kubectl-logs.png)  
  
### Create a Service
Ekspos *Pod* ke publik dengan menggunakan perintah berikut:  
```
$ kubectl expose deployment hello-node --type=LoadBalancer --port=808
```
Perintah di atas akan membuat sebuah *Service* yang bertujuan untuk membuat aplikasi dapat diakses dari jaringan luar.  
![membuat sebuah service dengan perintah kubectl expose deployment](./10-kubectl-expose-deployment.png)  

Jalankan perintah `kubectl get services` untuk melihat *Service* yang sudah dibuat.  
![melihat service dengan perintah kubectl get services](./11-kubectl-get-services.png)  
  
Jalankan perintah berikut:  
```
$ sudo minikube service hello-node
```
Perintah di atas akan menampilkan respon dari aplikasi yang sedang dijalankan tadi.  
![menjalankan perintah minikube service](./12-minikube-service.png)  
![tampilan respon aplikasi pada halaman web browser](./13-app-response.png)  
  
### Enable addons
Jalankan perintah `minikube addons list` untuk melihat daftar addons yang tersedia.  
![melihat daftar addons dengan perintah minikube addons list](./14-minikube-addons-list.png)  
  
Aktifkan sebuah addons, contohnya `metrics-server` dengan menggunakan perintah berikut:  
```
$ minikube addons enable metrics-server
```

Jalankan perintah berikut:  
```
$ kubectl get pod,svc -n kube-system
```
![melihat pod dengan perintah kubectl get pod](./15-kubectl-get-pod.png)  

Nonaktifkan addons `metrics-server` dengan perintah:  
```
$ minikube addons disable metrics-server
```
  
### Cleaning up
Hapus *Deployment* dan *Service* yang sudah dibuat sebelumnya.  
```
$ kubectl delete service hello-node
$ kubectl delete deployment hello-node
```
![menghapus service dan deployment dengan perintah kubectl delete](./16-kubectl-delete.png)  
  
Hentikan minikube dengan perintah:  
```
$ minikube stop
```  
![menghentikan minikube dengan perintah minikube stop](./17-minikube-stop.png)  
  
Apabila sudah tidak ingin menggunakan minikube lagi, dapat dilakukan penghapusan dengan perintah:  
```
$ minikube delete
```
  
Selesai.