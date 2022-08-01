package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"sync"
	"time"

	v13 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const numIngresses = 500

func main() {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := kubernetes.NewForConfigOrDie(config)

	ns := v13.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: "ingresses",
		},
	}

	fmt.Println("Creating ns..")
	_, err = client.CoreV1().Namespaces().Create(context.Background(), &ns, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating whoami..")
	home := path.Join(os.Getenv("HOME"), "/.kube/config")
	cmd := exec.Command("kubectl", "--kubeconfig", home, "apply", "-f", "./manifests/1-whoami.yaml")

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	/*
		fmt.Println("Creating special ingress..")
		ing := makeIngress(100, ns.Name)
		for {
			if _, err := client.NetworkingV1().Ingresses(ns.Name).Create(context.Background(), ing, v1.CreateOptions{}); err == nil {
				break
			}
			time.Sleep(time.Second)
			println("Retrying...")
		}
		return
	*/

	rand.Seed(time.Now().Unix())

	go func() {
		for {
			// pods, err := client.CoreV1().Pods(ns.Name).List(context.Background(), v1.ListOptions{})
			// if err != nil {
			// 	panic(err)
			// }
			//
			// end := rand.Intn(len(pods.Items))
			// fmt.Println("Deleting ", end, "/", len(pods.Items))
			//
			// for i := 0; i < end; i++ {
			// 	err := client.CoreV1().Pods(ns.Name).Delete(context.Background(), pods.Items[i].Name, v1.DeleteOptions{})
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// }

			fmt.Println("Deleting pods...")
			for {
				if err := client.CoreV1().Pods(ns.Name).DeleteCollection(context.Background(), v1.DeleteOptions{}, v1.ListOptions{}); err == nil {
					break
				}
				time.Sleep(time.Second)
				println("Retrying...")
			}

			time.Sleep(30 * time.Second)
		}
	}()

	time.Sleep(20 * time.Second)

	for {

		fmt.Println("Creating half of ingresses...")

		for i := 0; i < numIngresses/2; i++ {
			ing := makeIngress(i, ns.Name)
			for {
				if _, err := client.NetworkingV1().Ingresses(ns.Name).Create(context.Background(), ing, v1.CreateOptions{}); err == nil {
					break
				}
				time.Sleep(time.Second)
				println("Retrying...")
			}
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("Sending requests...")
			for i := 0; i < numIngresses; i++ {
				req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
				if err != nil {
					panic(err)
				}
				req.Host = fmt.Sprintf("toto.%d.localhost", i)

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					res.Body.Close()
					continue
					// panic(err)
				}
				res.Body.Close()

				if res.StatusCode != 200 {
					fmt.Println("Request failed for: ", req.Host, "->", res.StatusCode)
				}

				time.Sleep(100 * time.Millisecond)
			}

		}()

		fmt.Println("Creating second half ingresses...")

		for i := numIngresses / 2; i < numIngresses; i++ {
			ing := makeIngress(i, ns.Name)
			for {
				if _, err := client.NetworkingV1().Ingresses(ns.Name).Create(context.Background(), ing, v1.CreateOptions{}); err == nil {
					break
				}
				time.Sleep(time.Second)
				println("Retrying...")
			}
		}
		println("INGRESSES DONE")

		time.Sleep(time.Minute)
		// time.Sleep(5 * time.Second)

		fmt.Println("Deleting ingresses...")

		for {
			if err := client.NetworkingV1().Ingresses(ns.Name).DeleteCollection(context.Background(), v1.DeleteOptions{}, v1.ListOptions{}); err == nil {
				break
			}
			time.Sleep(time.Second)
			println("Retrying...")
		}

		// time.Sleep(30 * time.Second)
		wg.Wait()
	}
}

func makeIngress(n int, ns string) *v12.Ingress {
	return &v12.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:      "ingress-" + strconv.Itoa(n),
			Namespace: ns,
			Labels: map[string]string{
				"foo": "bar",
			},
		},
		Spec: v12.IngressSpec{
			Rules: []v12.IngressRule{
				{
					Host: "*." + strconv.Itoa(n) + ".localhost",
					IngressRuleValue: v12.IngressRuleValue{
						HTTP: &v12.HTTPIngressRuleValue{
							Paths: []v12.HTTPIngressPath{{
								Path:     "/",
								PathType: stringPtr("Prefix"),
								Backend: v12.IngressBackend{
									Service: &v12.IngressServiceBackend{
										Name: "whoami",
										Port: v12.ServiceBackendPort{
											Number: 80,
										},
									},
								},
							}},
						},
					},
				},
			},
		},
	}
}

func stringPtr(s string) *v12.PathType {
	p := v12.PathType(s)
	return &p
}
