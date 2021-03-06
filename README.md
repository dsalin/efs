# EFS (Group Based Encrypted Filesystem)

EFS is an encrypted Linux filesystem. Primary difference of EFS is that all files are encrypted group based, not file based.

# How to Use

Currently, EFS is being unstable and not ready to be tested. However, it is useful to view the source code to understand
the logic and general expected functionality.

# Dependencies

EFS is highly dependent on Bazil Fuse implementation: https://github.com/bazil/fuse
All other dependencies are standard Golang libraries.

# EFS CLI Interface

Currently, the only way to interact with EFS is through its command line interface, called efsctl. In this section, we are going to explore the basic functionality with examples.

### 1. Initialize a filesystem as a directory:

`$ efsctl init -s test.dir -t /mnt -k ~/.ssh/id_rsa`

Here, we initialize and mound the EFS on folder `test.dir` (next version of EFS will mount a file instead) and mounting it to `/mnt` path, giving our RSA key for initial setup. Using this key, EFS creates a `default` group and assigns this RSA key to it. Therefore, any file manipulations will use this key for encryption. In addition to that, EFS stores the hash of this key, so that next time when someone mounts the filesystem, it is only possible to read the real contents of this group using this exact key. EFS will take a SHA256 of the provided key and check it with the group hash key used before during initialization.

### 2. Create work group with different key:

`$ efsctl create group work -k ~/.ssh/work`

This time, we use another group with a new key. Therefore, a new entry in the EFS is created with the hashed name of this group (the EFS directory listing will be shown in the end of this section).

### 3. Do regular file work:

`$ echo “Work doc” > worksheet.txt`

Using regular Linux commands to create/manipulate files. At this stage, only small files are considered.

### 4. Switch to previously created group:

`$ efsctl switch personal -k ~/.ssh/personal`

Now, we are switching to another group with another key provided. At this point, a SHA256 of this key will be taken and compared with a corresponding SHA256 of the key used during creation of this group. In case of mismatch, an error will occur.

### 4. Unmount current filesystem, releasing all temporary info:

`$ efsctl unmount -s test.dir`

# Project Status

EFS is in active development. Unfortunately, not all mentioned functionality is currently working. However, a good foundation is laid for all of steps described in previous section: core data structures, encryption, FUSE (Bazil) bindings, command line interface structure.
