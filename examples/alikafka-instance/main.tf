resource "alicloud_alikafka_instance" "default" {
  name              = "${var.name}"
  topic_quota       = "${var.topic_quota}"
  disk_type         = "${var.disk_type}"
  disk_size         = "${var.disk_size}"
  deploy_type       = "${var.deploy_type}"
  io_max            = "${var.io_max}"
  eip_max           = "${var.eip_max}"
  vpc_id            = "${var.vpc_id}"
  vswitch_id        = "${var.vswitch_id}"
  zone_id           = "${var.zone_id}"
}