**Friday January 15, 2021**

### 1. The first day of git course :git commom commands ###

**Initialize git**

```git
git init
```

**check the status of files under the folder**

```git
git status
```

**File that need to be managed by git**

```git
git add filename  #for example:git add index.html
```

**Manage all files under he current folder**

```git
git add .
```

**Use git to generate the first managed version**

```git
git commit -m 'v1'  #this version is described in single quotation marks
```

**After modifyig any file under the folder managed by git,you can use git add to hand over the file to git for management.**

```git
git add .
git commit -m "v2"     #Generate the second version of the project.
```

**Friday  January  15,2021**

#  The second course

**Error:**   #Personal information configuraton,user name,mailbox,tell git who is generating this version control.

**View version record**

```git
git log
```

**Three regions of git** 

**git common commands**

1. Clone code 

   ```git
   http address: git clone https://github.com/gitzreal93/dev-tester.git
   ```

   ```git
   git adrress: git clone git@github.com:GitDzreal93/dev-tester.git
   ```

   

2. Pull substitution code 

   ```git
   git pull
   ```

   

3. switch branch

   ```git
   git checkout branch_name
   ```

   

4. View workspace status

   ```git
   git status
   ```

   

5. View past submission records

   ```git
   git log
   ```

   

6. Submit to staging area

   ```git
   git add
   ```

7. Submit to the local warehouse

   ```git
   git commit -m "submission comment"
   git branch -M main
   git remote add origin https://github.com/ceoco/godemo.git
   git push -u origin main
   ```

   

8. Submit to remote warehouse

   ```git
   git push
   ```

   

9.Delete staging area file

```git
rm filename.file extension
git rm filename.file extension
```

 # 3.github introduction and project search #

**common comand **

1. project name find

   ```git
   in:name dev-tester
   ```

   

2. project description find

   ```git
   in:description
   ```

   

3. through REAdme description find

   ```git
   in:readme python
   ```

   

4. sigh up stars 

   ```git
   stars>1000
   ```

   

5. sight up fork number

   ```git
   forks:>500
   ```

   

6. project languager

   ```git
   language:java
   ```

   

7. project auther find

   ```git
   user：GitDzreal93
   ```

   

8. project size find

   ```git
   size:>=500
   ```

   

9. multi finding

   ```git
   in:name test  in:description in:readme pyton stars:>50
   ```

   

# 4.github download acceleration

# 5.github projet project creation and upload