# QNAP NAS: Samba (SMB): Enable Windows ACL Support (Fixed a Bug on the NAS)

### Note

This fixes a bug found in the model `QNAP TS-469 Pro` with firmware version `4.2.0 Build 20160130`.

I expect it to work on other models with firmware versions `4.2.x` as well.

In short, the required module was installed in a wrong location, but not noticed by QNAP.

Copying the module to the correct location solves the problem.

To jump directly to the solution, click [here](#solution).



### Enabling Windows ACL Support on the NAS

To enable Windows ACL support, follow these steps on the NAS:

1. `Control Panel > Privilege Settings > Shared Folders > Advanced Permissions`;
1. Check the checkbox `Enable Windows ACL support`;
1. Click the `Apply All` button.

However, after making the change, the shares could not be reached anymore.

To jump directly to the solution, click [here](#solution).



### Attempt to Connect to Samba on the NAS from Mac OS X with `Finder`

To connect the share with `Finder`, follow these steps on the Mac:

1. `Finder > Go > Connect to Server...`;
1. Enter `smb://username@host/share`, with `username`, `host`, and `share` replaced by your actual values;
1. Click the `Connect` button;
1. Enter the correct password for the `username`;
1. Click the `Connect` button.

However, the following error is shown:

> **There was a problem connecting to the server “host”.**
>
> The share does not exist on the server.
> Please check the share name, and then try again.

To jump directly to the solution, click [here](#solution).



### Attempt to Connect to Samba on the NAS from Mac OS X with `smbutil`

1. Start `Terminal`;
1. Enter the following command, with `username`, and `host` replaced by your actual values:

       smbutil view //username@host

1. Enter the correct password for the `username`;
1. Press the `Enter` key on your keyboard.

However, the following error is shown:

>     smbutil: share name doesn't exist: No such file or directory

To jump directly to the solution, click [here](#solution).



### Find the Solution (Advanced, May Skip)

1. Start `Terminal`;
1. Connect to the NAS with `ssh` with the following command, with `host` replaced by your actual value:

       ssh admin@host

1. Examine the command line arguments of the `smbd` process:

       ps | grep smbd

   Here is the response:

   >      9880 admin       532 S   grep smbd
   >     27686 admin      3088 S   /usr/local/samba/sbin/smbd -l /var/log -D -s /etc/config/smb.conf
   >     27697 admin      1340 S   /usr/local/samba/sbin/smbd -l /var/log -D -s /etc/config/smb.conf

   So, it is known that:

   1. The `smbd` executable is located at `/usr/local/samba/sbin/smbd`;
   1. The log files are located in the directory `/var/log`;
   1. The config file is located at `/etc/config/smb.conf`.
1. List the directory `/var/log`:

       ls -al /var/log

   Here is the response:

   >     drwxr-xr-x    5 admin    administ       420 Mar 16 11:01 ./
   >     drwxrwxrwx   10 admin    administ       560 Mar 16 10:56 ../
   >     -rw-r--r--    1 admin    administ     96487 Mar 16 10:59 SDMD.log
   >     drwx------    4 admin    administ        80 Mar 15 21:00 cores/
   >     drwxr-xr-x    2 admin    guest          100 Mar 15 23:46 cups/
   >     -rw-r--r--    1 admin    administ      2095 Mar 15 21:00 hal_app.log
   >     ----------    1 admin    administ     10043 Mar 16 09:58 hal_daemon.log
   >     -rw-r--r--    1 admin    administ      4881 Mar 16 11:01 hal_lib.log
   >     -rw-r--r--    1 admin    administ     10285 Mar 16 11:01 hal_lib.log.bak
   >     -rw-r--r--    1 admin    administ       268 Mar 15 20:59 hal_util_net.log
   >     -rw-r--r--    1 admin    administ       275 Mar 16 10:44 log.nmbd
   >     -rw-r--r--    1 admin    administ         0 Mar 16 03:50 log.nmbd.old
   >     -rw-r--r--    1 admin    administ      3378 Mar 16 10:23 log.smbd
   >     -rw-r--r--    1 admin    administ     10878 Mar 16 10:21 log.smbd.old
   >     -rw-r--r--    1 admin    administ         0 Mar 15 21:00 messages
   >     drwxr-xr-x    2 admin    administ       140 Mar 15 21:00 network/
   >     -rw-------    1 admin    administ       333 Mar 15 21:00 php-fpm.log
   >     lrwxrwxrwx    1 admin    administ        20 Mar 16 04:57 samba -> /usr/local/samba/var/
   >     -rw-r--r--    1 admin    administ    435123 Mar 16 11:01 storage_lib.log
   >     -rw-r--r--    1 admin    administ   1048716 Mar 16 10:08 storage_lib.log.bak
   >     -rw-r--r--    1 admin    administ      2356 Mar 15 20:57 storage_util.log

   So, the log file is expected to be `/var/log/log.smbd`;
1. Read the log file:

       cat /var/log/log.smbd

   Here is tail of the response (`xuserx` is the username used):

   >     [2016/03/16 03:50:47.447463,  1] ../librpc/ndr/ndr.c:247(ndr_print_debug)
   >            negotiate: struct NEGOTIATE_MESSAGE
   >               Signature                : 'NTLMSSP'
   >               MessageType              : NtLmNegotiate (1)
   >               NegotiateFlags           : 0x62888215 (1653113365)
   >                      1: NTLMSSP_NEGOTIATE_UNICODE
   >                      0: NTLMSSP_NEGOTIATE_OEM
   >                      1: NTLMSSP_REQUEST_TARGET
   >                      1: NTLMSSP_NEGOTIATE_SIGN
   >                      0: NTLMSSP_NEGOTIATE_SEAL
   >                      0: NTLMSSP_NEGOTIATE_DATAGRAM
   >                      0: NTLMSSP_NEGOTIATE_LM_KEY
   >                      0: NTLMSSP_NEGOTIATE_NETWARE
   >                      1: NTLMSSP_NEGOTIATE_NTLM
   >                      0: NTLMSSP_NEGOTIATE_NT_ONLY
   >                      0: NTLMSSP_ANONYMOUS
   >                      0: NTLMSSP_NEGOTIATE_OEM_DOMAIN_SUPPLIED
   >                      0: NTLMSSP_NEGOTIATE_OEM_WORKSTATION_SUPPLIED
   >                      0: NTLMSSP_NEGOTIATE_THIS_IS_LOCAL_CALL
   >                      1: NTLMSSP_NEGOTIATE_ALWAYS_SIGN
   >                      0: NTLMSSP_TARGET_TYPE_DOMAIN
   >                      0: NTLMSSP_TARGET_TYPE_SERVER
   >                      0: NTLMSSP_TARGET_TYPE_SHARE
   >                      1: NTLMSSP_NEGOTIATE_EXTENDED_SESSIONSECURITY
   >                      0: NTLMSSP_NEGOTIATE_IDENTIFY
   >                      0: NTLMSSP_REQUEST_NON_NT_SESSION_KEY
   >                      1: NTLMSSP_NEGOTIATE_TARGET_INFO
   >                      1: NTLMSSP_NEGOTIATE_VERSION
   >                      1: NTLMSSP_NEGOTIATE_128
   >                      1: NTLMSSP_NEGOTIATE_KEY_EXCH
   >                      0: NTLMSSP_NEGOTIATE_56
   >               DomainNameLen            : 0x0000 (0)
   >               DomainNameMaxLen         : 0x0000 (0)
   >               DomainName               : NULL
   >               WorkstationLen           : 0x0000 (0)
   >               WorkstationMaxLen        : 0x0000 (0)
   >               Workstation              : NULL
   >               Version: struct ntlmssp_VERSION
   >                   ProductMajorVersion      : NTLMSSP_WINDOWS_MAJOR_VERSION_6 (6)
   >                   ProductMinorVersion      : NTLMSSP_WINDOWS_MINOR_VERSION_1 (1)
   >                   ProductBuild             : 0x1db0 (7600)
   >                   Reserved: ARRAY(3)
   >                       [0]                      : 0x0f (15)
   >                       [1]                      : 0x00 (0)
   >                       [2]                      : 0x00 (0)
   >                   NTLMRevisionCurrent      : UNKNOWN_ENUM_VALUE (0)
   >     [2016/03/16 03:50:47.451606,  2] auth/auth.c:309(check_ntlm_password)
   >       check_ntlm_password:  authentication for user [xuserx] -> [xuserx] -> [xuserx] succeeded
   >     [2016/03/16 03:50:47.474831,  0] smbd/vfs.c:173(vfs_init_custom)
   >       error probing vfs module 'acl_xattr': NT_STATUS_UNSUCCESSFUL
   >     [2016/03/16 03:50:47.474943,  0] smbd/vfs.c:315(smbd_vfs_init)
   >       smbd_vfs_init: vfs_init_custom failed for acl_xattr
   >     [2016/03/16 03:50:47.475008,  0] smbd/service.c:902(make_connection_snum)
   >       vfs_init failed for service IPC$
   >     [2016/03/16 03:50:47.475826,  0] smbd/vfs.c:173(vfs_init_custom)
   >       error probing vfs module 'acl_xattr': NT_STATUS_UNSUCCESSFUL
   >     [2016/03/16 03:50:47.475940,  0] smbd/vfs.c:315(smbd_vfs_init)
   >       smbd_vfs_init: vfs_init_custom failed for acl_xattr
   >     [2016/03/16 03:50:47.476005,  0] smbd/service.c:902(make_connection_snum)
   >       vfs_init failed for service IPC$
   >     [2016/03/16 03:50:47.477022,  2] smbd/smb2_server.c:2917(smbd_smb2_request_incoming)
   >       smbd_smb2_request_incoming: client read error NT_STATUS_PIPE_BROKEN

   So, it is known that the problem is due to that the module `'acl_xattr'` could not be loaded;
1. Examine the build flags of `smbd`:

       /usr/local/samba/sbin/smbd -b

   Here is the response:

   >     Build environment:
   >        Built by:    root@BuildServer-105
   >        Built on:    Sat Jan 30 14:51:18 CST 2016
   >        Built using: gcc
   >        Build host:  Linux BuildServer-105 2.6.22-14-server #1 SMP Sun Oct 14 23:34:23 GMT 2007 i686 GNU/Linux
   >        SRCDIR:      /root/daily_build/4.2.0/DataService/CIFS/samba-3.6.18/source3
   >        BUILDDIR:    /root/daily_build/4.2.0/DataService/CIFS/samba-3.6.18/source3
   >
   >     Paths:
   >        SBINDIR: /usr/local/samba/sbin
   >        BINDIR: /usr/local/samba/bin
   >        SWATDIR: /usr/local/samba/swat
   >        CONFIGFILE: /etc/config/smb.conf
   >        LOGFILEBASE: /usr/local/samba/var
   >        LMHOSTSFILE: /etc/config/lmhosts
   >        LIBDIR: /usr/local/samba/lib
   >        MODULESDIR: /usr/local/samba/lib
   >        SHLIBEXT: so
   >        LOCKDIR: /usr/local/samba/var/locks
   >        STATEDIR: /usr/local/samba/var/locks
   >        CACHEDIR: /usr/local/samba/var/locks
   >        PIDDIR: /usr/local/samba/var/locks
   >        SMB_PASSWD_FILE: /usr/local/samba/private/smbpasswd
   >        PRIVATE_DIR: /usr/local/samba/private
   >        NCALRPCDIR: /usr/local/samba/var/ncalrpc
   >        NMBDSOCKETDIR: /usr/local/samba/var/nmbd
   >
   >      System Headers:
   >        HAVE_SYS_ACL_H
   >        ... (skipped) ...
   >        HAVE_SYS_STATVFS_H
   >        ... (skipped) ...
   >        HAVE_SYS_VFS_H
   >        HAVE_SYS_WAIT_H
   >        HAVE_SYS_XATTR_H
   >
   >      Headers:
   >        HAVE_ACL_LIBACL_H
   >        ... (skipped) ...
   >
   >      UTMP Options:
   >        ... (skipped) ...
   >
   >      HAVE_* Defines:
   >        ... (skipped) ...
   >        HAVE_FGETXATTR
   >        HAVE_FLISTXATTR
   >        ... (skipped) ...
   >        HAVE_FREMOVEXATTR
   >        ... (skipped) ...
   >        HAVE_FSETXATTR
   >        ... (skipped) ...
   >        HAVE_GETXATTR
   >        ... (skipped) ...
   >        HAVE_LGETXATTR
   >        ... (skipped) ...
   >        HAVE_LISTXATTR
   >        HAVE_LLISTXATTR
   >        ... (skipped) ...
   >        HAVE_LREMOVEXATTR
   >        ... (skipped) ...
   >        HAVE_LSETXATTR
   >        ... (skipped) ...
   >        HAVE_POSIX_ACLS
   >        ... (skipped) ...
   >        HAVE_REMOVEXATTR
   >        ... (skipped) ...
   >        HAVE_SETXATTR
   >        ... (skipped) ...
   >        HAVE_STATVFS_F_FLAG
   >        ... (skipped) ...
   >
   >      --with Options:
   >        ... (skipped) ...
   >
   >      Build Options:
   >        ... (skipped) ...
   >        STAT_STATVFS64
   >        ... (skipped) ...
   >        static_decl_vfs
   >        ... (skipped) ...
   >        static_init_vfs
   >        vfs_acl_tdb_init
   >        vfs_acl_xattr_init
   >        ... (skipped) ...
   >        vfs_streams_xattr_init
   >        ... (skipped) ...
   >        vfs_xattr_tdb_init
   >
   >     Type sizes:
   >        sizeof(char):         1
   >        sizeof(int):          4
   >        sizeof(long):         4
   >        sizeof(long long):    8
   >        sizeof(uint8):        1
   >        sizeof(uint16):       2
   >        sizeof(uint32):       4
   >        sizeof(short):        2
   >        sizeof(void*):        4
   >        sizeof(size_t):       4
   >        sizeof(off_t):        8
   >        sizeof(ino_t):        8
   >        sizeof(dev_t):        8
   >
   >     Builtin modules:
   >         pdb_ldap pdb_smbpasswd pdb_tdbsam pdb_wbc_sam idmap_ldap idmap_tdb idmap_passdb idmap_nss idmap_rid nss_info_template auth_sam auth_unix auth_winbind auth_wbc auth_server auth_domain auth_builtin vfs_default vfs_posixacl

   So, it is known that:

   1. The `Samba` package has ACL support;
   1. The ACL support is not a builtin module, but a separated `vfs` module;
   1. The modules are located at `/usr/local/samba/lib`.
1. List the directory `/usr/local/samba/lib`:

       ls -al /usr/local/samba/lib

   Here is the response:

   >     drwxr-xr-x    3 admin    administ      4096 Mar 16 03:28 ./
   >     drwxrwxr-x    7 admin    administ      4096 Feb 28  2006 ../
   >     -rw-r--r--    1 admin    administ     31052 Jan 30 16:31 acl_xattr.so
   >     -rw-r--r--    1 admin    administ    131072 Jan 30 16:31 lowcase.dat
   >     lrwxrwxrwx    1 admin    administ        20 Mar 15 20:58 smb.conf -> /etc/config/smb.conf
   >     -rw-r--r--    1 admin    administ    131072 Jan 30 16:31 upcase.dat
   >     -rw-r--r--    1 admin    administ     65536 Jan 30 16:31 valid.dat
   >     drwxr-xr-x    2 admin    administ      4096 Mar 16 10:12 vfs/

   Note that `acl_xattr.so` is found here!

1. List the directory `/usr/local/samba/lib/vfs`:

       ls -al /usr/local/samba/lib/vfs

   Here is the response:

   >     drwxr-xr-x    2 admin    administ      4096 Mar 16 10:12 ./
   >     drwxr-xr-x    3 admin    administ      4096 Mar 16 03:28 ../
   >     -rwxr-xr-x    1 admin    administ     18572 Jan 30 16:31 aio_fork.so*
   >     -rwxr-xr-x    1 admin    administ     18668 Jan 30 16:31 aio_pthread.so*
   >     -rwxr-xr-x    1 admin    administ     10240 Jan 30 16:31 dirsort.so*

   At this point, it is easy to spot the solution!



### Solution

1. Start `Terminal` (if not done so);
1. Connect to the NAS with `ssh` with the following command, with `host` replaced by your actual value (if not done so):

       ssh admin@host

1. Copy the required module to the correct directory:

       cp /usr/local/samba/lib/acl_xattr.so /usr/local/samba/lib/vfs/

1. Correct the permission of the module:

       chmod 755 /usr/local/samba/lib/vfs/acl_xattr.so

1. Logout the `ssh` session:

       exit

1. Retry connecting to Samba.
