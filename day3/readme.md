# DevOps Fundamentals — Day 3 (First Practical Session)

This session focused on enhancing Bash scripting skills by building a **Demo Real-Time Log Rotation & Cleanup Scheduler**.

## Project Overview

The project simulates a **production-like log management system**, performing:

- Continuous generation of logs
- Automatic rotation when logs exceed a size threshold
- Compression of old logs
- Cleanup of logs older than a defined retention period

This demonstrates the **real-world relevance of automation and monitoring** in system administration and DevOps workflows.

## Why This Project Matters in DevOps

- **Log management is critical** to prevent disk overflows and maintain application stability.  
- **Automation reduces manual intervention and errors**, ensuring reliable operations.  
- **Bash scripting forms the foundation** for system automation and DevOps tasks.  

## Key Technical Concepts Demonstrated

- Bash **loops, conditionals, and functions** for automation
- Safe file handling: directory creation, file existence checks, and error prevention
- Use of Linux system commands:  
  - `du` → checking file sizes  
  - `find` → locating old logs  
  - `mv` → rotating logs  
  - `gzip` → compressing logs  
- Real-time scheduling using a loop to monitor logs continuously
- Handling edge cases such as missing directories or empty variables

## Challenges and Solutions

| Challenge | Solution |
|-----------|---------|
| Log directory didn’t exist initially | Added directory creation checks |
| Disk usage and log size variables could be empty | Added checks before performing arithmetic comparisons |
| Simulating log growth for testing rotation | Wrote a loop to append lines to the log continuously |

## Key Takeaways

- Real-time monitoring and event-driven automation are essential skills in DevOps.  
- Bash scripting enables **rapid prototyping of system management solutions**.  
- Practical projects like this provide experience in handling **production-level challenges** like log rotation, disk monitoring, and cleanup automation.  

## Real-World Applications

- Automating log rotation and cleanup on servers  
- Monitoring disk usage to prevent outages  
- Event-driven automation in system administration  
- Foundational knowledge for CI/CD pipelines and production monitoring

## Project Repository

The full project and scripts can be found here:  
[GitHub Link]  

---

This basic project showcases **real-world use cases of shell scripting** and demonstrates how simple automation can have a significant impact in DevOps workflows.

