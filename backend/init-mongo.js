
db.createUser(
 {
   user:"rust",
   pwd: "123",
   roles: [
     { role:"readWrite", db:"service_3d_db" },
     { role:"userAdminAnyDatabase", db:"admin" },
     { role:"dbAdminAnyDatabase", db:"admin" },
     { role:"readWriteAnyDatabase", db:"admin" }
   ]	
 }
)
