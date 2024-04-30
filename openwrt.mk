include $(TOPDIR)/rules.mk

PKG_NAME:=mediamtx
PKG_VERSION:=v0.0.0
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/noranetworks/kvds
PKG_SOURCE_VERSION:=$(PKG_VERSION)

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_USE_MIPS16:=0

GO_PKG:=github.com/noranetworks/kvds
GO_PKG_LDFLAGS_X:=github.com/noranetworks/kvds/internal/core.version=$(PKG_VERSION)

include $(INCLUDE_DIR)/package.mk
include $(TOPDIR)/feeds/packages/lang/golang/golang-package.mk

GO_MOD_ARGS:=-buildvcs=false

define Package/mediamtx
  SECTION:=net
  CATEGORY:=Network
  TITLE:=kvds
  URL:=https://github.com/noranetworks/kvds
  DEPENDS:=$(GO_ARCH_DEPENDS)
endef

define Package/kvds/description
  ready-to-use server and proxy that allows users to publish, read and proxy live video and audio streams through various protocols
endef

$(eval $(call GoBinPackage,kvds))
$(eval $(call BuildPackage,kvds))
