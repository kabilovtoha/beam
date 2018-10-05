// Copyright 2018 The Beam Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "utility/io/buffer.h"

namespace beam {

class Node;
class HttpMsgCreator;

namespace explorer {

/// node->explorer adapter interface
struct IAdapter {
    using Ptr = std::unique_ptr<IAdapter>;

    virtual ~IAdapter() = default;

    /// Returns body for /status request
    virtual const io::SharedBuffer& get_status(HttpMsgCreator& packer) = 0;

    virtual const io::SharedBuffer& get_block(HttpMsgCreator& packer, uint64_t height) = 0;

    virtual void get_blocks(HttpMsgCreator& packer, io::SerializedMsg& out, uint64_t startHeight, uint64_t endHeight) = 0;
};

IAdapter::Ptr create_adapter(Node& node);

}} //namespaces
