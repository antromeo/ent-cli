# ent-cli

## MoSCoW

### Mo (Must have) 
- [x] quickstart with minikube and tekton pipeline crd
- [ ] ent bundle with podman -> replace with binary version obtained with pkg
- [x] clients to manage all Entando CRs
- [ ] provide executables 
- [ ] migrates the commands
- [ ] modularity
- [ ] strict versioning
- [ ] uses ko for build container images -> To handle several platform sha

### S (Should have)
- [x] cluster native, interaction with the Entando CRs
- [ ] a command for entando backup
- [ ] set cpu/ram pod limits
- [ ] set http/https
- [x] set DB type(with entando-clients) 
- [ ] set DB internal or external
- [ ] set internal or external keycloak 

### Co (Could have)
- [x] profile management
- [ ] extensibility of plugins in a kubectl-style

### W (Won't have)
- [ ] external dependencies
