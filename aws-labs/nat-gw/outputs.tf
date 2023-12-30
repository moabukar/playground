# Output the subnet IDs for ref
output "db_subnet_ids" {
  value = aws_subnet.db_subnet.*.id
}

output "app_subnet_ids" {
 value = aws_subnet.app_subnet.*.id
}

output "web_subnet_ids" {
 value = aws_subnet.web_subnet.*.id
}

# Outputs for route table IDs
output "app_route_table_ids" {
  value = aws_route_table.app_route_table.*.id
}

output "db_route_table_ids" {
  value = aws_route_table.db_route_table.*.id
}
