#%%
from faker import Faker
import hashlib
from pathlib import Path

#%%
def generate_users(n:int, fake:Faker)-> list[dict]:
    users:list=[]
    
    for _ in range(0, n, 1):
        user:dict={
            "name": fake.name()
            , "dob": fake.date_of_birth(minimum_age=12, maximum_age=70).strftime("%Y-%m-%d")
            , "phone_number": fake.phone_number()
            , "email": fake.email()
            , "password":fake.password()
            , "active": True
            , "created_at": fake.date_time().strftime("%Y-%m-%d %H:%M:%S")
        }
        users.append(user)
        # 
    
    return users

#%%
if __name__=="__main__":
    fake:Faker=Faker()
    users:list=generate_users(n=1000, fake=fake)
    
    with open(file=Path.joinpath(Path(__name__).parent, "output", "user.txt"), mode="w+") as write_io:
        for user in users:
            write_io.write(f'''('{user["name"]}', '{user["dob"]}', '{user["phone_number"]}', '{user["email"]}', '{user["password"]}', {user["active"]}, '{user["created_at"]}'),\n''')