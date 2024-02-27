resource "ovh_cloud_project_kube" "my_kube_cluster" {
  name         = "my_kube_cluster_cdktf"
  region       = "GRA7"
}

#resource "ovh_cloud_project_kube_nodepool" "my_pool" {
#  kube_id       = ovh_cloud_project_kube.my_kube_cluster[count.index].id
#  name          = "my-pool" //Warning: "_" char is not allowed!
#  flavor_name   = "b2-7"
#  desired_nodes = 1
#  max_nodes     = 1
#  min_nodes     = 1
#}

# output "kubeconfig_file" {
#   value     = ovh_cloud_project_kube.my_kube_cluster.kubeconfig
#   sensitive = true
# }

#resource "local_file" "kubeconfig" {
#    content     = ovh_cloud_project_kube.my_kube_cluster[count.index].kubeconfig
#    filename = "my-kube-cluster-cdktf.yml"
#}