# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class HomebrewUrlo < Formula
  desc "A simple CLI tool to open URLs from the command line"
  homepage "https://github.com/ryo034/homebrew-urlo"
  version "1.0.27"
  license "Apache-2.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/ryo034/homebrew-urlo/releases/download/v1.0.27/homebrew-urlo_Darwin_x86_64.tar.gz"
      sha256 "c20d0355eefebd4fb01ffcadb74427ff40695676ec7a8888fd6be3627564a4ea"

      def install
        bin.install "urlo"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/ryo034/homebrew-urlo/releases/download/v1.0.27/homebrew-urlo_Darwin_arm64.tar.gz"
      sha256 "c895abb296aa94df048d14a3c45fe71a595a6b34dd998846f0145402684bedc0"

      def install
        bin.install "urlo"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/ryo034/homebrew-urlo/releases/download/v1.0.27/homebrew-urlo_Linux_arm64.tar.gz"
      sha256 "2547645740fdbb0ce8f2e82cb37b06638dae6d10eb992436c1d413d8b1493c47"

      def install
        bin.install "urlo"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ryo034/homebrew-urlo/releases/download/v1.0.27/homebrew-urlo_Linux_x86_64.tar.gz"
      sha256 "462d34e0503f097c32282f290ecdc8b904dc8a33c3ef76cc04a796153bfbfbe4"

      def install
        bin.install "urlo"
      end
    end
  end

  test do
    urlo --version
  end
end
